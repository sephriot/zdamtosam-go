package db

import (
	"database/sql"
	"zdamtosam.pl/src/model"
)

func GetUsers(db *sql.DB) []model.User {
	rows, err := db.Query("SELECT * FROM users;")
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var users []model.User
	for rows.Next() {
		var user model.User
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Picture, &user.TutorId)
		users = append(users, user)
	}

	return users
}
