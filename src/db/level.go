package db

import (
	"database/sql"
	"zdamtosam/src/model"
)

func GetLevels(db *sql.DB) []model.Level {
	rows, err := db.Query("SELECT * FROM levels;")
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var levels []model.Level
	for rows.Next() {
		var level model.Level
		rows.Scan(&level.Id, &level.Name)
		levels = append(levels, level)
	}

	return levels
}
