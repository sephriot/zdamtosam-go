package frontend

import (
	"html/template"
	"net/http"
	"regexp"
	"zdamtosam.pl/src/db"
	"zdamtosam.pl/src/frontend/tmplengine"
)

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	pageRegex := regexp.MustCompile(`page=[0-9]+&?`)

	data := h.prepareTemplateData(r)
	data["Exercises"] = db.GetExercisesBySearchQuery(h.db, r.URL.Query().Get("query"))
	data["RawQuery"] = template.HTMLAttr(pageRegex.ReplaceAllString(r.URL.RawQuery, ""))
	data["PageTitle"] = "ZdamToSam | Wyszukaj zadanie"

	tmplengine.Render(w, data,
		tmplengine.FS_PATH_PREFIX+"templates/search.html",
		tmplengine.FS_PATH_PREFIX+"templates/navbar.html",
		tmplengine.FS_PATH_PREFIX+"templates/categories.html",
		tmplengine.FS_PATH_PREFIX+"templates/subcategories.html",
		tmplengine.FS_PATH_PREFIX+"templates/exercises.html",
		tmplengine.FS_PATH_PREFIX+"templates/exercise.html")
}
