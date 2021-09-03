package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../database"
	"github.com/gorilla/mux"
)

type Data struct {
	Name             string
	Email            string
	Password         string
	Confirm_Password string
}

type Teacher struct {
	Id       int
	Email    string
	Password string
	Name     string
	Phone    string
	Status   bool
}

type Status struct {
	Status bool
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var data Data

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "something when wrong with parser body: %v", err)
		return
	}
	var isExists Teacher

	database.DB.Where("email = ?", data.Email).First(&isExists)

	if isExists.Email != "" {
		fmt.Fprintf(w, "email exists!")
		return
	}
	if data.Password == data.Confirm_Password {
		fmt.Fprintf(w, "password and confirm password do not match")
		return
	}

	teacher := Teacher{
		Email:    data.Email,
		Password: data.Password,
		Name:     data.Name,
	}
	database.DB.Create(&teacher)
	fmt.Fprintf(w, "completed create %v", teacher)
}

func GetTeachers(w http.ResponseWriter, r *http.Request) {
	var allTeachers []Teacher

	database.DB.Find(&allTeachers)
	json.NewEncoder(w).Encode(allTeachers)
}

func GetTeacherById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var teacher Teacher
	database.DB.Where("id = ?", id).First(&teacher)
	if teacher.Id == 0 {
		fmt.Fprintf(w, "not found teacher have id: %v", id)
		return
	}

	json.NewEncoder(w).Encode(&teacher)

}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var teacher Teacher
	database.DB.Where("id = ?", id).First(&teacher)
	if teacher.Id == 0 {
		fmt.Fprintf(w, "not found teacher have id: %v", teacher.Id)
		return
	}

	var data Teacher
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
	}

	teacher.Password = data.Password
	teacher.Name = data.Name
	teacher.Phone = data.Phone

	database.DB.Model(&teacher).Updates(teacher)
	fmt.Fprintf(w, "updated teacher have id: %v", teacher.Id)
}

func SetStatusTeacher(w http.ResponseWriter, r *http.Request) {

	var status Status

	if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var teacher Teacher
	database.DB.Where("id = ?", id).First(&teacher)
	if teacher.Id == 0 {
		fmt.Fprintf(w, "not foung teacher have id: %v", id)
	}

	database.DB.Model(&teacher).Update("status", status.Status)
	fmt.Fprintf(w, "set status teacher have id: %v completed", id)
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	teacher := Teacher{
		Id: id,
	}
	database.DB.Delete(&teacher)
	fmt.Fprintf(w, "deleted teacher have have id: %v", id)
}

