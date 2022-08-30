package frontend

import (
	"net/http"
	"regexp"
	"zdamtosam/src/db"
	"zdamtosam/src/frontend/tmplengine"
)

func extractLevelPath(path string) string {
	firstPath := regexp.MustCompile("/level/[1-9]+")
	level := firstPath.FindStringSubmatch(path)
	if level != nil {
		return level[0]
	}
	return ""
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	levels := db.GetLevels(h.db)
	categories := db.GetCategories(h.db)
	levelPath := extractLevelPath(r.URL.Path)
	data := map[string]interface{}{
		"Levels":      levels,
		"Categories":  categories,
		"CurrentPath": r.URL.Path,
		"LevelPath":   levelPath,
	}
	tmplengine.Render(w, data, "templates/index.html", "templates/navbar.html", "templates/categories.html")
}
