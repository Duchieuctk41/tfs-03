package server

import (
	"fmt"
	"net/http"

	// import local package
	"../calc"
)

func RunServer() {
	fmt.Println("Server is listening in http://localhost:3000")
	// Defer function will be called when process exists.
	defer func() {
		fmt.Println("Server is Stopped")
	}()

	http.HandleFunc("/calculator", calc.Calculator)

	//run server
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic("Error when running server")
	}


}
