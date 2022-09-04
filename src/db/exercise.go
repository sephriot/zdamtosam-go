package db

import (
	"database/sql"
	"math"
	"sort"
	"strings"
	"zdamtosam/src/model"
)

func unique(input []model.Exercise) []model.Exercise {
	u := make([]model.Exercise, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val.Id]; !ok {
			m[val.Id] = true
			u = append(u, val)
		}
	}

	return u
}

func sortBestMarch(input []model.Exercise, searchSubQueries []string) []model.Exercise {
	type ExerciseWithScore struct {
		exercise model.Exercise
		score    int
	}
	scored := make([]ExerciseWithScore, 0, len(input))
	for _, exercise := range input {
		value := 0
		lastPos := -1
		for _, subQuery := range searchSubQueries {
			nowPos := strings.Index(exercise.Task, subQuery)
			if nowPos != -1 {
				value++
			}
			if lastPos != -1 && lastPos < nowPos {
				value++
			}
			lastPos = nowPos
		}
		scored = append(scored, ExerciseWithScore{score: value, exercise: exercise})
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].score > scored[j].score
	})
	ret := make([]model.Exercise, 0, len(input))
	for _, s := range scored {
		ret = append(ret, s.exercise)
	}
	return ret
}

func GetExercisesBySubcategoryId(db *sql.DB, subcategoryId string) []model.Exercise {
	rows, err := db.Query("SELECT id, task, date FROM exercises WHERE subcategory_id = ? AND visible = 1;", subcategoryId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var exercises []model.Exercise
	for rows.Next() {
		var exercise model.Exercise
		rows.Scan(&exercise.Id, &exercise.Task, &exercise.Date)
		exercises = append(exercises, exercise)
	}
	return exercises
}

func GetExercisesBySearchQuery(db *sql.DB, searchQuery string) []model.Exercise {

	split := strings.Split(searchQuery, " ")
	// Consider up to 25 first keywords
	split = split[:int(math.Min(float64(len(split)), 25))]
	exercises := make([]model.Exercise, 0)
	for _, subQuery := range split {
		query := "SELECT exercises.id AS id, task, subcategory_id, level_id, category_id FROM `exercises` JOIN subcategories ON exercises.subcategory_id = subcategories.id JOIN categories on categories.id = subcategories.category_id  WHERE task LIKE ? LIMIT 20;"
		rows, err := db.Query(query, "%"+subQuery+"%")
		if err != nil {
			panic(err)
		}
		var exercise model.Exercise
		for rows.Next() {
			rows.Scan(&exercise.Id, &exercise.Task, &exercise.SubcategoryId, &exercise.LevelId, &exercise.CategoryId)
			exercises = append(exercises, exercise)
		}
		_ = rows.Close()
	}

	exercises = unique(exercises)
	exercises = sortBestMarch(exercises, split)
	exercises = exercises[:int(math.Min(float64(len(exercises)), 50))] // return 20 best matches

	return exercises
}

func GetWrongAnswersForExerciseById(db *sql.DB, id string) []string {
	query := "SELECT value FROM wrong_answers WHERE exercise_id = ?;"
	rows, err := db.Query(query, id)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	answers := make([]string, 0)
	for rows.Next() {
		answer := ""
		rows.Scan(&answer)
		answers = append(answers, answer)
	}
	return answers
}

func GetPreviousExerciseId(db *sql.DB, exercise model.Exercise) int {
	query := "SELECT id FROM exercises WHERE id < ? AND subcategory_id = ? AND level_id = ? AND visible = 1 ORDER BY id DESC LIMIT 1;"
	rows, err := db.Query(query, exercise.Id, exercise.SubcategoryId, exercise.LevelId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	id := 0
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}

func GetNextExerciseId(db *sql.DB, exercise model.Exercise) int {
	query := "SELECT id FROM exercises WHERE id > ? AND subcategory_id = ? AND level_id = ? AND visible = 1 ORDER BY id LIMIT 1;"
	rows, err := db.Query(query, exercise.Id, exercise.SubcategoryId, exercise.LevelId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	id := 0
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}

func GetExerciseById(db *sql.DB, id string) model.Exercise {
	query := "SELECT exercises.id, task, hint, stepByStep, image, level_id, subcategory_id, answer" +
		" FROM (SELECT * FROM exercises WHERE id = ?  LIMIT 1) AS exercises JOIN" +
		" (SELECT value as answer, exercise_id FROM correct_answers) AS answers" +
		" ON exercises.id = answers.exercise_id LIMIT 1;"

	rows, err := db.Query(query, id)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var exercise model.Exercise
	for rows.Next() {
		rows.Scan(&exercise.Id, &exercise.Task, &exercise.Hint, &exercise.StepByStep,
			&exercise.Image, &exercise.LevelId, &exercise.SubcategoryId, &exercise.Answer)
	}
	exercise.Options = append(GetWrongAnswersForExerciseById(db, id), exercise.Answer)
	exercise.NextId = GetNextExerciseId(db, exercise)
	exercise.PreviousId = GetPreviousExerciseId(db, exercise)

	return exercise
}
