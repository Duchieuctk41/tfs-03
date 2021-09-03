package controllers

import (
	"fmt"
	"net/http"
	"time"
)

func CreateCalendar(w http.ResponseWriter, r *http.Request) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"

	// Calling Parse() method with its parameters
	tm, _ := time.Parse(layout, "Feb 4, 2014 at 6:05pm (PST)")

	// Returns output
	fmt.Println(tm)
	fmt.Fprintf(w, "hello world")

	// var isExists Calendar

	// database.DB.Where("teacher_id = ?", calendar.Teacher_Id).Where("time_start = ?", calendar.Time_Start).First(&isExists)

	// if isExists.Id != 0 {
	// 	fmt.Fprintf(w, "calendar exists!")
	// 	return
	// }

	// database.DB.Create(&calendar)
	// fmt.Fprintf(w, "completed create %v for teacher_id: %v", calendar.Id, calendar.Teacher_Id)
}
