package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	// app.LoadEnvironment()
	// db := db.GetDB()
	// webDir := os.Getenv("APP_WEB_DIR")
	// driver, _ := mysql.WithInstance(db, &mysql.Config{})
	// m, _ := migrate.NewWithDatabaseInstance("file://"+webDir+"/sql/migrations", "mysql", driver)

	// m.Steps(2)
}
