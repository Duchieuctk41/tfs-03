package routes

import (
	"fmt"
	"log"
	"net/http"

	"../controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/calc", controllers.Calculator).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	fmt.Println("server is running in http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handler))

}
