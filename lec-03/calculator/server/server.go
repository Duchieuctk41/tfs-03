package server

import (
	"fmt"
	"net/http"

	"../calc"
)

func serveFiles(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p == "/" {
		fmt.Println(p)
		p = "./client"
	}
	http.ServeFile(w, r, p)
}

func RunServer() {
	fmt.Println("Server is listening in http://localhost:3000")
	//Defer function will be called process exitst.
	defer func() {
		fmt.Println("Server is Stopped")
	}()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./client"))))
	http.HandleFunc("/calc", calc.Calculator)
	// run server
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic("Error when running server")
	}
}
