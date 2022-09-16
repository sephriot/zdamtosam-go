package backend

import (
	"database/sql"
	"net/http"
	"strings"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/login") {
		return
	}

	w.Write([]byte(""))
}
