package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"../database"
	"github.com/gorilla/mux"
)

type ParseCalendar struct {
	Teacher_Id int
	Class_Id   string
	Time_Start string
	Time_End   string
}

type Calendar struct {
	Id         int
	Teacher_Id int
	Class_Id   string
	Time_Start time.Time
	Time_End   time.Time
}

func CreateCalendar(w http.ResponseWriter, r *http.Request) {

	var parseCalendar ParseCalendar
	if err := json.NewDecoder(r.Body).Decode(&parseCalendar); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
		return
	}

	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	ts, _ := time.Parse(layout, parseCalendar.Time_Start)
	te, _ := time.Parse(layout, parseCalendar.Time_End)

	var calendar Calendar
	calendar.Teacher_Id = parseCalendar.Teacher_Id
	calendar.Class_Id = parseCalendar.Class_Id
	calendar.Time_Start = ts
	calendar.Time_End = te

	var isExists Calendar

	database.DB.Where("teacher_id = ?", calendar.Teacher_Id).Where("time_start = ?", calendar.Time_Start).First(&isExists)

	if isExists.Id != 0 {
		fmt.Fprintf(w, "calendar exists!")
		return
	}

	database.DB.Create(&calendar)
	fmt.Fprintf(w, "completed create for teacher_id: %v", calendar.Teacher_Id)
}

func GetCalendarByClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	class_id := vars["class_id"]

	var allCalendars []Calendar

	database.DB.Where("class_id = ?", class_id).First(&allCalendars)
	json.NewEncoder(w).Encode(allCalendars)
}

func GetCalendarByTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teacher_id := vars["teacher_id"]

	var allCalendars []Calendar

	database.DB.Where("teacher_id = ?", teacher_id).First(&allCalendars)
	json.NewEncoder(w).Encode(allCalendars)
}

func UpdateCalendar(w http.ResponseWriter, r *http.Request) {

	var data Calendar
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "something when wrong while parser body: %v", err)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]

	var calendar Calendar

	database.DB.Where("teacher_id = ?", id).First(&calendar)
	if calendar.Id == 0 {
		fmt.Fprintf(w, "not found calendar id: %v", id)
		return
	}

	data.Id = calendar.Id

	database.DB.Model(&calendar).Updates(data)
	fmt.Fprintf(w, "updated calendar id: %v", id)
}

func DeleteCalendar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	calendar := Calendar{
		Id: id,
	}

	database.DB.Delete(&calendar)
	fmt.Fprintf(w, "Deleted class have id: %v", id)
}
