package server

import (
	"../calc"
	"fmt"
	"net/http"
)

func RunServer() {
	fmt.Println("Server is listening in http://localhost:3000")
	//Defer function will be called process exitst.
	defer func() {
		fmt.Println("Server is Stopped")
	}()
		
	http.HandleFunc("/calc", calc.Calculator)

	// run server
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic("Error when running server")
	}
}
