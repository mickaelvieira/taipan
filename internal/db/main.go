package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// GetDB returns a database connection
func GetDB() *sql.DB {
	if db == nil {
		fmt.Println("init DB")
		dsn := os.Getenv("APP_DB_USER") + ":" + os.Getenv("APP_DB_PWD") + "@tcp(" + os.Getenv("APP_DB_ADDR") + ")/" + os.Getenv("APP_DB_NAME")
		params := "parseTime=true"
		var err error
		db, err = sql.Open("mysql", dsn+"?"+params)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("reuse DB")
	}

	return db
}
