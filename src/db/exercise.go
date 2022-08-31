package db

import (
	"database/sql"
	"zdamtosam/src/model"
)

func GetExercisesBySubcategoryId(db *sql.DB, subcategoryId string) []model.Exercise {
	rows, err := db.Query("SELECT id, task FROM exercises WHERE subcategory_id = ? AND visible = 1;", subcategoryId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var exercises []model.Exercise
	for rows.Next() {
		var exercise model.Exercise
		rows.Scan(&exercise.Id, &exercise.Task)
		exercises = append(exercises, exercise)
	}
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
	exercise.Options = GetWrongAnswersForExerciseById(db, id)
	exercise.NextId = GetNextExerciseId(db, exercise)
	exercise.PreviousId = GetPreviousExerciseId(db, exercise)

	return exercise
}
