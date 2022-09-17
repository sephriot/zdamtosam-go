package frontend

import (
	"database/sql"
	"firebase.google.com/go/v4/auth"
	"net/http"
	zdamtosamDB "zdamtosam.pl/src/db"
)

type Handler struct {
	db        *sql.DB
	auth      *auth.Client
	userCache *zdamtosamDB.UserCache
}

func NewHandler(db *sql.DB, auth *auth.Client) *Handler {
	return &Handler{db, auth, zdamtosamDB.NewUserCache()}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	h.Index(w, r)
	return

}
