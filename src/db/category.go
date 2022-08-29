package db

import (
	"database/sql"
	"zdamtosam/src/model"
)

func GetCategories(db *sql.DB) []model.Category {
	rows, err := db.Query("SELECT * FROM categories;")
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}

	return categories
}
