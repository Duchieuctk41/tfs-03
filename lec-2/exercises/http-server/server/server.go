package server

import (
	"fmt"
	"net/http"
	handlers2 "tfs-03/lec-2/exercises/http-server/handlers"
)

func RunServer() {
	fmt.Println("Server listening in http://localhost:3000/")
	// Defer function will be called when process exists.
	defer func() {
		fmt.Println("Server is stopped.")
	}()

	http.HandleFunc("/", handlers2.HomePage)

	http.HandleFunc("/hello", handlers2.Hello)

	http.HandleFunc("/api/student", handlers2.GetAStudent)

	http.HandleFunc("/api/students", handlers2.GetStudents)

	//run server
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic("Error when running server")
	}
}
