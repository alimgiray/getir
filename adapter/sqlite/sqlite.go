package sqlite

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(models ...interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
	}
	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("failed to migrate database: %s", err.Error())
	}
	return db
}
