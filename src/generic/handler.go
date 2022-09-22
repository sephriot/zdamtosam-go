package generic

import (
	"database/sql"
	"firebase.google.com/go/v4/auth"
	"log"
	"net/http"
	zdamtosamDB "zdamtosam.pl/src/db"
	"zdamtosam.pl/src/model"
)

type Handler struct {
	Db        *sql.DB
	Auth      *auth.Client
	UserCache *zdamtosamDB.UserCache
}

func (h *Handler) GetLoggedUser(r *http.Request) model.User {
	cookie, err := r.Cookie("__session")
	var ret model.User
	if err != nil {
		log.Default().Println(err)
		return ret
	}

	if cookie.Value == "" {
		return ret
	}

	cachedToken := h.UserCache.Get(cookie.Value)
	var userRecord *auth.UserRecord
	if cachedToken == nil {
		token, err := zdamtosamDB.VerifyIDToken(h.Auth, cookie.Value)
		if err != nil {
			log.Default().Println(err)
			return ret
		}

		userRecord, err = zdamtosamDB.GetUser(h.Auth, token.UID)
		if err != nil {
			log.Default().Println(err)
			return ret
		}
		h.UserCache.Put(cookie.Value, token, userRecord)
		cachedToken = h.UserCache.Get(cookie.Value)
	}

	ret.Id = cachedToken.UID
	ret.Email = cachedToken.Email
	ret.Picture = cachedToken.PhotoURL
	ret.Name = cachedToken.DisplayName

	return ret
}
