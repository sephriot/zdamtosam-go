package backend

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
	w.Write([]byte(""))
}
