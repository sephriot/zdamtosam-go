package frontend

import (
	"net/http"
	"zdamtosam/src/db"
	"zdamtosam/src/frontend/tmplengine"
	"zdamtosam/src/model"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	levels := db.GetLevels(h.db)
	pathParams := getPathParams(r.URL.Path)
	var categories []model.Category
	levelPath := ""
	if pathParams["level"] != "" {
		levelPath = "/level/" + pathParams["level"]
		categories = db.GetCategoriesByLevel(h.db, pathParams["level"])
	} else {
		categories = db.GetCategories(h.db)
	}

	data := map[string]interface{}{
		"Levels":      levels,
		"Categories":  categories,
		"CurrentPath": r.URL.Path,
		"LevelPath":   levelPath,
		"Breadcrumbs": getBreadcrumbs(h.db, r.URL.Path),
	}
	tmplengine.Render(w, data, "templates/index.html", "templates/navbar.html", "templates/categories.html")
}
