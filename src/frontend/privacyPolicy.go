package frontend

import (
	"net/http"
	"zdamtosam/src/db"
	"zdamtosam/src/frontend/tmplengine"
)

func (h *Handler) PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	levels := db.GetLevels(h.db)
	levelPath := ""
	pageTitle := "ZdamToSam | Warunki korzystania z serwisu"
	pageDescription := "Polityka prywatności serwisu ZdamToSam"

	data := map[string]interface{}{
		"Levels":          levels,
		"LevelPath":       levelPath,
		"CurrentPath":     r.URL.Path,
		"Breadcrumbs":     getBreadcrumbs(h.db, r.URL.Path),
		"PageTitle":       pageTitle,
		"PageDescription": pageDescription,
	}
	tmplengine.Render(w, data, "templates/privacy-policy.html", "templates/navbar.html")
}
