package backend

import (
	"encoding/xml"
	"net/http"
	"zdamtosam/src/db"
)

func (h *Handler) Sitemap(w http.ResponseWriter, r *http.Request) {
	xml.NewEncoder(w).Encode(db.GetUsers(h.db))
}
