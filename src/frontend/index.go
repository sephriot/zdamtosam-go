package frontend

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"zdamtosam.pl/src/db"
	"zdamtosam.pl/src/frontend/tmplengine"
	"zdamtosam.pl/src/model"
)

func rotate(s []string, k int) []string {
	if k < 0 || len(s) == 0 {
		return s
	}
	r := len(s) - k%len(s)
	s = append(s[r:], s[:r]...)

	return s
}

func getCurrentLevelName(levels []model.Level, currentId string) string {
	for _, l := range levels {
		stringId := strconv.Itoa(l.Id)
		if stringId == currentId {
			return l.Name
		}
	}
	return ""
}

func getCurrentCategoryName(categories []model.Category, currentId string) string {
	for _, c := range categories {
		stringId := strconv.Itoa(c.Id)
		if stringId == currentId {
			return c.Name
		}
	}
	return ""
}

func getCurrentSubcategoryName(subcategories []model.Subcategory, currentId string) string {
	for _, s := range subcategories {
		stringId := strconv.Itoa(s.Id)
		if stringId == currentId {
			return s.Name
		}
	}
	return ""
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
	pageTitle := "ZdamToSam"
	pageDescription := "Zadania z matmy na każdym poziomie. Tutaj znajdziesz zadania, podpowiedzi i pełne rozwiązania. " +
		"Ucz się samodzielnie lub z korepetytorem. Śledź swoje postępy, a na pewno zdasz na 5."
	if pathParams["level"] != "" {
		levelPath = "/level/" + pathParams["level"]
		levelName := getCurrentLevelName(levels, pathParams["level"])
		pageTitle = "ZdamToSam | " + levelName
		pageDescription = "Zadania dla poziomu " + strings.ToLower(levelName)
		categories = db.GetCategoriesByLevel(h.db, pathParams["level"])
	} else {
		categories = db.GetCategories(h.db)
	}
	if pathParams["category"] != "" {
		categoryPath = levelPath + "/category/" + pathParams["category"]
		subcategories = db.GetSubcategories(h.db, pathParams["category"])
		categoryName := getCurrentCategoryName(categories, pathParams["category"])
		pageTitle = "ZdamToSam | " + categoryName
		pageDescription = "Zadania z działu " + strings.ToLower(categoryName)
	}
	if pathParams["subcategory"] != "" {
		subcategoryPath = categoryPath + "/subcategory/" + pathParams["subcategory"]
		exercises = db.GetExercisesBySubcategoryId(h.db, pathParams["subcategory"])
		subcategoryName := getCurrentSubcategoryName(subcategories, pathParams["subcategory"])
		pageTitle = "ZdamToSam | " + subcategoryName
		pageDescription = "Zadania z rodziału " + strings.ToLower(subcategoryName)
	}
	if pathParams["exercise"] != "" {
		exercisePath = subcategoryPath + "/exercise/" + pathParams["exercise"]
		exercise = db.GetExerciseById(h.db, pathParams["exercise"])
		rotateBy := rand.Intn(4)
		answerIndex = (3 + rotateBy) % 4
		exercise.Options = rotate(exercise.Options, rotateBy)
		pageTitle = strings.ReplaceAll(exercise.Task, "`", "")
		pageDescription = strings.ReplaceAll(exercise.Task, "`", "")
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
		"PageTitle":       pageTitle,
		"PageDescription": pageDescription,
	}
	tmplengine.Render(w, data, tmplengine.TEMPLATE_PATH_PREFIX+"templates/index.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/navbar.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/homepage.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/categories.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/subcategories.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/exercises.html",
		tmplengine.TEMPLATE_PATH_PREFIX+"templates/exercise.html")
}
