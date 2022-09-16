package frontend

import (
	"database/sql"
	"firebase.google.com/go/v4/auth"
	"net/http"
)

type Handler struct {
	db   *sql.DB
	auth *auth.Client
}

func NewHandler(db *sql.DB, auth *auth.Client) *Handler {
	return &Handler{db, auth}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	h.Index(w, r)
	return

}
