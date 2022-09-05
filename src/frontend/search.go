package frontend

import (
	"html/template"
	"net/http"
	"regexp"
	"zdamtosam.pl/src/db"
	"zdamtosam.pl/src/frontend/tmplengine"
	"zdamtosam.pl/src/model"
)

// Search TODO: deduplicate
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	levels := db.GetLevels(h.db)
	var categories []model.Category
	var subcategories []model.Subcategory
	var exercise model.Exercise
	levelPath := ""
	categoryPath := ""
	subcategoryPath := ""
	exercisePath := ""
	answerIndex := 0
	exercises := db.GetExercisesBySearchQuery(h.db, r.URL.Query().Get("query"))
	pageRegex := regexp.MustCompile(`page=[0-9]+&?`)
	rawQuery := template.HTMLAttr(pageRegex.ReplaceAllString(r.URL.RawQuery, ""))

	pageTitle := "ZdamToSam | Wyszukaj zadanie"
	pageDescription := "Zadania z matmy na każdym poziomie. Tutaj znajdziesz zadania, podpowiedzi i pełne rozwiązania. Ucz się samodzielnie lub z korepetytorem. Śledź swoje postępy, a na pewno zdasz na 5."

	data := map[string]interface{}{
		"Levels":          levels,
		"Categories":      categories,
		"Subcategories":   subcategories,
		"Exercises":       exercises,
		"Exercise":        exercise,
		"AnswerIndex":     answerIndex,
		"CurrentPath":     r.URL.Path,
		"LevelPath":       levelPath,
		"CategoryPath":    categoryPath,
		"SubcategoryPath": subcategoryPath,
		"ExercisePath":    exercisePath,
		"Breadcrumbs":     getBreadcrumbs(h.db, r.URL.Path),
		"QueryPage":       r.URL.Query().Get("page"),
		"RawQuery":        rawQuery,
		"PageTitle":       pageTitle,
		"PageDescription": pageDescription,
	}
	tmplengine.Render(w, data,
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/search.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/navbar.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/categories.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/subcategories.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/exercises.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/exercise.html")
}
