package main

import (
	"fmt"
	"learn/driver"
	models "learn/model"
	"learn/repository/repoimpl"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "password"
	dbname   = "meo"
)

func main() {
	db := driver.Connect(host, port, user, password, dbname)

	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	userRepo := repoimpl.NewUserRepo(db.SQL)

	user1 := models.User{
		ID:     1,
		Name:   "Son Tung",
		Gender: "male",
		Email:  "mtp@gmail.com",
	}
	user2 := models.User{
		ID:     2,
		Name:   "Dan Truong",
		Gender: "male",
		Email:  "dante@gmail.com",
	}

	userRepo.Insert(user1)
	userRepo.Insert(user2)

	users, _ := userRepo.Select()

	for i := range users {
		fmt.Println(users[i])
	}
}
