package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func NewDatabaseClient() *sql.DB {

	DbHost := os.Getenv("DB_HOST")
	DbPass := os.Getenv("DB_PASS")
	DbUser := os.Getenv("DB_USER")

	db, err := sql.Open("mysql", DbUser+":"+DbPass+"@"+DbHost+"/zdamtosam")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
