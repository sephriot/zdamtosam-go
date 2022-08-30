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

func GetCategoriesByLevel(db *sql.DB, level string) []model.Category {
	rows, err := db.Query("SELECT DISTINCT categories.id, categories.name FROM `exercises` JOIN subcategories ON subcategories.id = exercises.subcategory_id JOIN categories ON categories.id = subcategories.category_id WHERE exercises.visible = 1 AND exercises.level_id = ? ORDER BY id;", level)
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

func GetCategoryNameById(db *sql.DB, id string) string {
	rows, err := db.Query("SELECT name FROM categories WHERE id = ?;", id)
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
