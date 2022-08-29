package frontend

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, ok := vars["section"]
	if !ok {
		h.Index(w, r)
		return
	}
}
