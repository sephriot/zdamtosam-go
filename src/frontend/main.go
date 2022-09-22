package frontend

import (
	"database/sql"
	"firebase.google.com/go/v4/auth"
	"net/http"
	zdamtosamDB "zdamtosam.pl/src/db"
	"zdamtosam.pl/src/generic"
)

type Handler struct {
	generic.Handler
}

func NewHandler(Db *sql.DB, Auth *auth.Client, UserCache *zdamtosamDB.UserCache) *Handler {
	return &Handler{generic.Handler{Db: Db, Auth: Auth, UserCache: UserCache}}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	h.Index(w, r)
	return

}
