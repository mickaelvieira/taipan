package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// GetDB returns a database connection
func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/taipan?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
