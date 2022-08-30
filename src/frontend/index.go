package frontend

import (
	"net/http"
	"regexp"
	"strings"
	"zdamtosam/src/db"
	"zdamtosam/src/frontend/tmplengine"
	"zdamtosam/src/model"
)

func getLevelPath(path string) string {
	firstPath := regexp.MustCompile("/level/[1-9]+")
	level := firstPath.FindStringSubmatch(path)
	if level != nil {
		return level[0]
	}
	return ""
}

func getLevel(path string) (string, bool) {
	s := strings.Split(getLevelPath(path), "/")
	if len(s) > 2 {
		return s[2], true
	}
	return "", false
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	levels := db.GetLevels(h.db)
	currentLevel, ok := getLevel(r.URL.Path)
	var categories []model.Category
	if ok {
		categories = db.GetCategoriesByLevel(h.db, currentLevel)
	} else {
		categories = db.GetCategories(h.db)
	}

	levelPath := getLevelPath(r.URL.Path)
	data := map[string]interface{}{
		"Levels":      levels,
		"Categories":  categories,
		"CurrentPath": r.URL.Path,
		"LevelPath":   levelPath,
	}
	tmplengine.Render(w, data, "templates/index.html", "templates/navbar.html", "templates/categories.html")
}
