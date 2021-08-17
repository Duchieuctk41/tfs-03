package routes

import (
	"log"
	"net/http"

	"../controllers"
	"../middlewares"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middlewares.ContentTypeChecking)

	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
	router.Use(middlewares.IsAuthenticated)

	router.HandleFunc("/", controllers.Home).Methods("GET")

	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products", controllers.AllsProduct).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
