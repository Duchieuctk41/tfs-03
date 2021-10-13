package main

import (
	"fmt"
	"hellogo/config"
	"hellogo/driver"
	"hellogo/handler"
	"net/http"
	// models "hellogo/model"
	// repoImpl "hellogo/repository/repoimpl"
)

func main() {
	driver.ConnectMongoDB(config.DB_USER, config.DB_PASS)

	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/user", handler.GetUser)

	fmt.Println("Server running localhost:3000")
	http.ListenAndServe(":3000", nil)
}

// userRepo := repoImpl.NewUserRepo(mongo.Client.Database(config.DB_NAME))

// user := models.User{
// 	Email:       "duchieu@gmail.com",
// 	Password:    "123456",
// 	DisplayName: "hieu hoc code",
// }

// err := userRepo.Insert(user)
// if err != nil {
// 	fmt.Println("insert ok")
// }

// user, _ := userRepo.FindUserEmail("duchieu@gmail.com")
// user, _ := userRepo.CheckLoginInfo("duchieu@gmail.com", "123456")
// fmt.Println(user)
