package db

import (
	"os"

	"github.com/jinzhu/gorm"

	// Import dialect for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Connect connects to database
func Connect() error {
	tmp, err := gorm.Open("postgres", os.Getenv("DB_URI"))

	if err != nil {
		return err
	}

	db = tmp
	return nil
}

// D returns database pointer
func D() *gorm.DB {
	return db
}

// Close closes connection to database
func Close() {
	if db != nil {
		db.Close()
	}
}
