package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../database"
	"github.com/gorilla/mux"
)

type Student struct {
	Id       int
	Email    string
	Password string
	Name     string
	Class_Id string
	Phone    string
	Status   bool
	Score    int
}

type ClassId struct {
	Class_Id string
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
		return
	}

	var student Student
	database.DB.Where("email = ?", data.Email).First(&student)
	if student.Email != "" {
		fmt.Fprintf(w, "email exists!")
		return
	}

	if data.Password != data.Confirm_Password {
		fmt.Fprintf(w, "password and password_confirm incorrect")
		return
	}

	student.Name = data.Name
	student.Email = data.Email
	student.Password = data.Password

	database.DB.Create(&student)
	fmt.Fprintf(w, "Create completed student")
}

func SetClassForStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var student Student
	database.DB.Where("id = ?", id).First(&student)
	if student.Id == 0 {
		fmt.Fprintf(w, "not found student have id: %v", id)
		return
	}

	var classId ClassId
	if err := json.NewDecoder(r.Body).Decode(&classId); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
		return
	}

	var class Class
	database.DB.Where("id = ?", classId.Class_Id).First(&class)
	if class.Id == "" {
		fmt.Fprintf(w, "not found class have id: %v", classId.Class_Id)
		return
	}

	database.DB.Model(&student).Update("class_id", classId.Class_Id)
	fmt.Fprintf(w, "set class_ id %v for student have id %v completed", classId.Class_Id, id)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	var allStudents []Student

	database.DB.Find(&allStudents)
	json.NewEncoder(w).Encode(allStudents)
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var student Student
	database.DB.Where("id = ?", id).First(&student)
	if student.Id == 0 {
		fmt.Fprintf(w, "not found student id: %v", id)
		return
	}

	json.NewEncoder(w).Encode(&student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var student Student
	database.DB.Where("id = ?", id).First(&student)
	if student.Id == 0 {
		fmt.Fprintf(w, "not found student id: %v", id)
		return
	}

	var data Student
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
	}

	student.Password = data.Password
	student.Name = data.Name
	student.Phone = data.Phone

	database.DB.Model(&student).Updates(student)
	fmt.Fprintf(w, "updated student have id: %v", student.Id)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	student := Student{
		Id: id,
	}
	database.DB.Delete(&student)
	fmt.Fprintf(w, "deleted student have have id: %v", id)
}

func SetScoreForStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var student Student
	database.DB.Where("id = ?", id).First(&student)
	if student.Id == 0 {
		fmt.Fprintf(w, "not found student have id: %v", id)
		return
	}

	type Score struct {
		Score int
	}
	var score Score
	if err := json.NewDecoder(r.Body).Decode(&score); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
		return
	}
	database.DB.Model(&student).Update("score", score.Score)
	fmt.Fprintf(w, "set score %v for student have id %v completed", score.Score, id)
}
