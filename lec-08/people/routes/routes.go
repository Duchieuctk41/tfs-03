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

	// class
	router.HandleFunc("/api/classes", controllers.CreateClass).Methods("POST")
	router.HandleFunc("/api/classes", controllers.GetClasses).Methods("GET")
	router.HandleFunc("/api/classes/{id}", controllers.GetClassById).Methods("GET")
	router.HandleFunc("/api/classes/{id}", controllers.UpdateClass).Methods("PUT")
	router.HandleFunc("/api/classes/{id}", controllers.DeleteClass).Methods("DELETE")

	// teacher
	router.HandleFunc("/api/teachers", controllers.CreateTeacher).Methods("POST")
	router.HandleFunc("/api/teachers", controllers.GetTeachers).Methods("GET")
	router.HandleFunc("/api/teachers/{id}", controllers.GetTeacherById).Methods("GET")
	router.HandleFunc("/api/teachers/{id}", controllers.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/api/teachers/{id}", controllers.SetStatusTeacher).Methods("PATCH")
	router.HandleFunc("/api/teachers/{id}", controllers.DeleteTeacher).Methods("DELETE")

	// student
	router.HandleFunc("/api/students", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/api/students", controllers.GetStudents).Methods("GET")
	router.HandleFunc("/api/students/{id}", controllers.GetStudentById).Methods("GET")
	router.HandleFunc("/api/students/{id}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/api/students/{id}", controllers.SetClassForStudent).Methods("PATCH")
	router.HandleFunc("/api/students/score/{id}", controllers.SetScoreForStudent).Methods("PATCH")
	router.HandleFunc("/api/students/{id}", controllers.DeleteStudent).Methods("DELETE")

	// calendar
	router.HandleFunc("/api/calendars", controllers.CreateCalendar).Methods("POST")
	// router.HandleFunc("/api/calendars/class/{class_id}", controllers.GetCalendarByClass).Methods("GET")
	// router.HandleFunc("/api/calendars/teacher/{teacher_id}", controllers.GetCalendarByTeacher).Methods("GET")
	// router.HandleFunc("/api/calendars/{id}", controllers.UpdateStudent).Methods("PUT")
	// router.HandleFunc("/api/calendars/{id}", controllers.DeleteStudent).Methods("DELETE")

	fmt.Println("server listening in http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
