package routes

import (
	"log"
	"net/http"

	"../controllers"
	"../middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":3000", handler))
}
