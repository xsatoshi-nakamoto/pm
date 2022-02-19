package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// the imports go here
var dbInstance *gorm.DB

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "host=localhost user=admin password=HTrang@091297 dbname=firmsone_pm port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(err)
	DB = db
	return DB
}
