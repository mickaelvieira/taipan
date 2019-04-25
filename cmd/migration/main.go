package main

import (
	"os"

	"github/mickaelvieira/taipan/internal/app"
	"github/mickaelvieira/taipan/internal/db"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	app.LoadEnvironment()
	db := db.GetDB()
	webDir := os.Getenv("APP_WEB_DIR")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance("file://"+webDir+"/sql/migrations", "mysql", driver)

	m.Steps(2)
}
