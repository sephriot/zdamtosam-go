package frontend

import (
	"net/http"
	"zdamtosam/src/db"
	"zdamtosam/src/frontend/tmplengine"
	"zdamtosam/src/model"
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
	}
	tmplengine.Render(w, data, "templates/search.html", "templates/navbar.html",
		"templates/categories.html", "templates/subcategories.html", "templates/exercises.html",
		"templates/exercise.html")
}
