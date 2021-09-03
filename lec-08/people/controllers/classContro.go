package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../database"
	"github.com/gorilla/mux"
)

type Class struct {
	Id          string
	Name        string
	Term        string
	Description string
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	var class Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
		return
	}

	var isExists Class
	database.DB.Where("id = ?", class.Id).First(&isExists)
	if isExists.Id != "" {
		fmt.Fprintf(w, "class exists!")
		return
	}

	database.DB.Create(&class)
	fmt.Fprintf(w, "created class have id: %v", class.Id)
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	var allClasses []Class

	database.DB.Find(&allClasses)
	json.NewEncoder(w).Encode(allClasses)
}

func GetClassById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var class Class
	database.DB.Where("id = ?", id).First(&class)
	if class.Id == "" {
		fmt.Fprintf(w, "not found class have id: %v", id)
		return
	}

	json.NewEncoder(w).Encode(&class)
}

func UpdateClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var class Class
	database.DB.Where("id = ?", id).First(&class)
	if class.Id == "" {
		fmt.Fprintf(w, "not found class have id: %v", id)
		return
	}

	var data Class
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
		return
	}

	class.Name = data.Name
	class.Term = data.Term
	class.Description = data.Description

	database.DB.Model(&class).Updates(class)
	fmt.Fprintf(w, "updated class have id: %v", id)
}

func DeleteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	class := Class{
		Id: id,
	}

	database.DB.Delete(&class)
	fmt.Fprintf(w, "Deleted class have id: %v", id)
}
