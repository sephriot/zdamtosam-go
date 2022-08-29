package backend

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"zdamtosam/src/db"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	// TODO: implement real API if needed
	json.NewEncoder(w).Encode(db.GetUsers(h.db))
}
