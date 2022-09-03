package frontend

import (
	"math/rand"
	"net/http"
	"zdamtosam/src/db"
	"zdamtosam/src/frontend/tmplengine"
	"zdamtosam/src/model"
)

func rotate(s []string, k int) []string {
	if k < 0 || len(s) == 0 {
		return s
	}
	r := len(s) - k%len(s)
	s = append(s[r:], s[:r]...)

	return s
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	levels := db.GetLevels(h.db)
	pathParams := getPathParams(r.URL.Path)
	var categories []model.Category
	var subcategories []model.Subcategory
	var exercises []model.Exercise
	var exercise model.Exercise
	levelPath := ""
	categoryPath := ""
	subcategoryPath := ""
	exercisePath := ""
	answerIndex := 0
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
	if pathParams["subcategory"] != "" {
		subcategoryPath = categoryPath + "/subcategory/" + pathParams["subcategory"]
		exercises = db.GetExercisesBySubcategoryId(h.db, pathParams["subcategory"])
	}
	if pathParams["exercise"] != "" {
		exercisePath = subcategoryPath + "/exercise/" + pathParams["exercise"]
		exercise = db.GetExerciseById(h.db, pathParams["exercise"])
		rotateBy := rand.Intn(4)
		answerIndex = (3 + rotateBy) % 4
		exercise.Options = rotate(exercise.Options, rotateBy)
	}

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
		"RawQuery":        "",
	}
	tmplengine.Render(w, data, "templates/index.html", "templates/navbar.html",
		"templates/categories.html", "templates/subcategories.html", "templates/exercises.html",
		"templates/exercise.html")
}
