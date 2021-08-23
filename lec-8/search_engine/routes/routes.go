package routes

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"../controllers"
)

func Init() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.GetHome).Methods("GET")
	router.HandleFunc("/search", controllers.SearchHandler).Methods("GET")

	fmt.Println("server is listening port: 3000")
	http.ListenAndServe(":3000", router)
}