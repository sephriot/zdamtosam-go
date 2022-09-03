package frontend

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/search" {
		h.Search(w, r)
		return
	}

	h.Index(w, r)
	return

}
