package frontend

import (
	"net/http"
	"zdamtosam.pl/src/db"
	"zdamtosam.pl/src/frontend/tmplengine"
)

func (h *Handler) PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	levels := db.GetLevels(h.Db)
	levelPath := ""
	pageTitle := "ZdamToSam | Polityka prywatności"
	pageDescription := "Polityka prywatności serwisu ZdamToSam"

	data := map[string]interface{}{
		"Levels":          levels,
		"LevelPath":       levelPath,
		"CurrentPath":     r.URL.Path,
		"Breadcrumbs":     getBreadcrumbs(h.Db, r.URL.Path),
		"PageTitle":       pageTitle,
		"PageDescription": pageDescription,
	}
	tmplengine.Render(w, data,
		tmplengine.FS_PATH_PREFIX+"templates/privacy-policy.html",
		tmplengine.FS_PATH_PREFIX+"templates/navbar.html")
}
