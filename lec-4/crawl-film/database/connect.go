package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:password@tcp/top_film"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database.")
	}
	DB = database
}
