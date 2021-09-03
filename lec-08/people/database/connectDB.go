package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() {
	database, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/people?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB = database
}
