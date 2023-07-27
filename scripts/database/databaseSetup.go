package database

import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

func DatabaseSetup() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
