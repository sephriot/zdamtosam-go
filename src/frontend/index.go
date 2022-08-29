package frontend

import (
	"net/http"
	"zdamtosam/src/db"
	"zdamtosam/src/frontend/tmplengine"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	levels := db.GetLevels(h.db)
	categories := db.GetCategories(h.db)
	data := map[string]interface{}{
		"Levels":     levels,
		"Categories": categories,
	}
	tmplengine.Render(w, data, "templates/index.html", "templates/navbar.html")
}
