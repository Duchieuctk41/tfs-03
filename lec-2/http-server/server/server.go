package server

import (
	"fmt"
	"net/http"

	//import local package
	"../handlers"
)

func RunServer() {
	fmt.Println("Server listening in http://localhost:3000/")
	// Defer function will be called when process exists.
	defer func() {
		fmt.Println("Server is stopped.")
	}()

	http.HandleFunc("/", handlers.HomePage)

	http.HandleFunc("/hello", handlers.Hello)

	http.HandleFunc("/api/student", handlers.GetAStudent)

	http.HandleFunc("/api/students", handlers.GetStudents)

	//run server
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic("Error when running server")
	}
}
