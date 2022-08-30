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
