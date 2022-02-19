package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// the imports go here
var db *gorm.DB

func InitDB() {
	dsn := "host=localhost user=admin password=HTrang@091297 dbname=firmsone_pm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err {
		fmt.Println("Cant connect db")
	}
}
