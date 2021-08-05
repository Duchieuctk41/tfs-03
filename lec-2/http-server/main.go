package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id int
	Name string
	Age int8
	Class []int
}

type Students []Student

func main() {

	fmt.Print("listening port: 3000")

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/api/student", GetAStudent)
	http.HandleFunc("/api/students", GetStudents)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello guys</h1>")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"msg": "connected About",
	}
	json.NewEncoder(w).Encode(data)
}

func GetAStudent(w http.ResponseWriter, r *http.Request) {
	var student = Student{1, "Hieu", 18, []int{1,2,3}}
	json.NewEncoder(w).Encode(student)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	var listStudent = Students{
		Student{1, "Hieu", 18, []int{1,2,3}},
		Student{1, "Hieu thu 2", 18, []int{1,2,3}},
	}
	json.NewEncoder(w).Encode(listStudent)
}
