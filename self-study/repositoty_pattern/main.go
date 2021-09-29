package main

import (
	"fmt"
	"learn/database"
	"learn/models"
	"learn/repo/repoimpl"
)

const (
	user     = "root"
	password = "password"
	host     = "localhost"
	port     = "3306"
	dbname   = "meo"
)

func main() {
	db := database.Connect(user, password, host, port, dbname)

	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}
	userRepo := repoimpl.NewUserRepo(db.SQL)

	user := models.User{
		ID:     12,
		Name:   "hieu chu nhat",
		Gender: "be de",
		Email:  "duchieu@gmail.com",
	}
	userRepo.Insert(user)

	users, _ := userRepo.Select()

	for i := range users {
		fmt.Println(users[i])
	}
}
