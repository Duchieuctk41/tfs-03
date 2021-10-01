package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const JsonContentType = "application/json"

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/login", ContentTypeCheck(login)).Methods("POST")

	fmt.Println("server is listening in localhost:3000")
	http.ListenAndServe(":3000", router)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"msg": "hello world",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func login(w http.ResponseWriter, r *http.Request) {
	// parse body
	type User struct {
		Email    string
		Password string
	}
	var data User
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "%s", err)
		return
	}

	// check password
	user := User{
		Email:    "hieuchunhat",
		Password: "hieuhoccode",
	}
	if data != user {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "email or password incorrect")
		return
	}

	// login success
	res := map[string]string{
		"msg": "login success",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ContentTypeCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != JsonContentType {
			fmt.Fprintf(w, "request only allow application/json")
			return
		}
		next.ServeHTTP(w, r)
	}
}
