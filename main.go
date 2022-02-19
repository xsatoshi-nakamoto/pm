package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=admin password=HTrang@091297 dbname=firmsone_pm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	var project = Project{ID: 5}
	db.Create(&project)
	fmt.Print(err)
	fmt.Print(db)
}
