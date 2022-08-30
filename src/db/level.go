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

func GetLevelNameById(db *sql.DB, id string) string {
	rows, err := db.Query("SELECT name FROM levels WHERE id = ?;", id)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var name string
		rows.Scan(&name)
		return name
	}
	return ""
}
