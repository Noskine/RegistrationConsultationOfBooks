package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
