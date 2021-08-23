package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../database"
	"../models"
)

type User struct {
	Email    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	req := User{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprintf(w, "error when parse body")
		return
	}
	var user models.User

	database.DB.Where("email = ?", req.Email).First(&user)
	if user.Id == 0 {
		fmt.Fprintf(w, "email not found")
		return
	}

	if user.Password != req.Password {
		fmt.Fprintf(w, "password incorrect!")
		return
	}

	cookie := http.Cookie{
		Name:    "logged_in",
		Value:   "true",
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 2), // 2 hours
	}

	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Hi %v, Have a good day", req.Email)

}
