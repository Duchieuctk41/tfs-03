package handlers

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	id int
	Name string
	Age int8
	class []int
}

type Students = []Student

func GetAStudent(w http.ResponseWriter, req *http.Request) {
	var student = Student{1, "Hieu", 18, []int{1,2,3}}
	json.NewEncoder(w).Encode(student)
}

func GetStudents(w http.ResponseWriter, req *http.Request) {
	var listStudents = Students{
		Student{1, "Hieu", 18, []int{1,2,3}},
		Student{2, "Hieu thu 2", 13, []int{5,6,7}},
	}
	json.NewEncoder(w).Encode(listStudents)
}
