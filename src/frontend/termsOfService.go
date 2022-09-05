package frontend

import (
	"net/http"
	"zdamtosam.pl/src/db"
	"zdamtosam.pl/src/frontend/tmplengine"
)

func (h *Handler) TermsOfService(w http.ResponseWriter, r *http.Request) {
	levels := db.GetLevels(h.db)
	levelPath := ""
	pageTitle := "ZdamToSam | Warunki korzystania z serwisu"
	pageDescription := "Warunki korzystania z serwisu ZdamToSam"

	data := map[string]interface{}{
		"Levels":          levels,
		"LevelPath":       levelPath,
		"CurrentPath":     r.URL.Path,
		"Breadcrumbs":     getBreadcrumbs(h.db, r.URL.Path),
		"PageTitle":       pageTitle,
		"PageDescription": pageDescription,
	}
	tmplengine.Render(w, data,
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/terms-of-service.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/navbar.html")
}
