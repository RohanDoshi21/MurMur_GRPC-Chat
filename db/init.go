package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var CONN *sql.DB

func InitDB() {
	db, err := sql.Open("postgres", "dbname=messagedb user=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	CONN = db
}
