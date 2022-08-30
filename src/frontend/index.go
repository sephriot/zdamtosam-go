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
	var subcategories []model.Subcategory
	levelPath := ""
	categoryPath := ""
	if pathParams["level"] != "" {
		levelPath = "/level/" + pathParams["level"]
		categories = db.GetCategoriesByLevel(h.db, pathParams["level"])
	} else {
		categories = db.GetCategories(h.db)
	}
	if pathParams["category"] != "" {
		categoryPath = levelPath + "/category/" + pathParams["category"]
		subcategories = db.GetSubcategories(h.db, pathParams["category"])
	}

	data := map[string]interface{}{
		"Levels":        levels,
		"Categories":    categories,
		"Subcategories": subcategories,
		"CurrentPath":   r.URL.Path,
		"LevelPath":     levelPath,
		"CategoryPath":  categoryPath,
		"Breadcrumbs":   getBreadcrumbs(h.db, r.URL.Path),
	}
	tmplengine.Render(w, data, "templates/index.html", "templates/navbar.html", "templates/categories.html", "templates/subcategories.html")
}
