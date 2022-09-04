package db

import (
	"database/sql"
	"zdamtosam.pl/src/model"
)

func GetSubcategories(db *sql.DB, categoryId string) []model.Subcategory {
	query := "SELECT subcategories.id, name FROM subcategories JOIN exercises ON subcategories.id = exercises.subcategory_id WHERE category_id = ? GROUP BY id ORDER BY subcategories.id;"
	rows, err := db.Query(query, categoryId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var subcategories []model.Subcategory
	for rows.Next() {
		var subcategory model.Subcategory
		rows.Scan(&subcategory.Id, &subcategory.Name)
		subcategories = append(subcategories, subcategory)
	}

	return subcategories
}

func GetSubcategoryNameById(db *sql.DB, id string) string {
	rows, err := db.Query("SELECT name FROM subcategories WHERE id = ?;", id)
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
