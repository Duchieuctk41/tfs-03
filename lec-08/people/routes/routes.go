package routes

import (
	"fmt"
	"log"
	"net/http"

	"../controllers"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter().StrictSlash(true)

	// home
	router.HandleFunc("/", controllers.Welcome).Methods("GET")

	// teacher
	router.HandleFunc("/api/teachers", controllers.CreateTeacher).Methods("POST")
	router.HandleFunc("/api/teachers", controllers.GetTeachers).Methods("GET")
	router.HandleFunc("/api/teachers/{id}", controllers.GetTeacher).Methods("GET")
	router.HandleFunc("/api/teachers/{id}", controllers.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/api/teachers/{id}", controllers.SetStatusTeacher).Methods("PATCH")
	router.HandleFunc("/api/teachers/{id}", controllers.DeleteTeacher).Methods("DELETE")

	fmt.Println("server listening in http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
