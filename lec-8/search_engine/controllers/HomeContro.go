package controllers

import (
	"log"
	"net/http"
	"html/template"
)
func GetHome(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("views/home.html")
		if err != nil {
			log.Print("Template parsing error: ", err)
		}
	
		err = t.Execute(w, nil)
		if err != nil {
			log.Print("Template executing error: ", err)
		}
}