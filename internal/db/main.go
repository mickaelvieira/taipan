package db

import (
	"database/sql"
	"github/mickaelvieira/taipan/internal/logger"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// GetDB returns a database connection
// NOTE: The MySQL driver convert all timestamp to UTC by default
func GetDB() *sql.DB {
	if db == nil {
		dsn := os.Getenv("APP_DB_USER") + ":" + os.Getenv("APP_DB_PWD") + "@tcp(" + os.Getenv("APP_DB_ADDR") + ")/" + os.Getenv("APP_DB_NAME")
		params := "parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_520_ci"
		var err error
		db, err = sql.Open("mysql", dsn+"?"+params)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		logger.Warn("reuse DB")
	}

	return db
}

// GetLastInsertID returns the last inserted ID as a string
func GetLastInsertID(r sql.Result) string {
	i, err := r.LastInsertId()
	if err != nil {
		return ""
	}
	return strconv.FormatInt(i, 10)
}
