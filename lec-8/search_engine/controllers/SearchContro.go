package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"../models"
)

type SearchResult struct {
	Students []models.Student `json:"students"`
	Input string `json:"input"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchInput := r.Form.Get("input")

	log.Print("Query database for: ", searchInput)

	students := models.SearchContent(searchInput)

	searchResult := SearchResult{
		Input: searchInput,
		Students: students,
	}

	jsonData, err := json.Marshal(searchResult)
	if err != nil {
		log.Print("JSON executing error: ", err)
		return
	}

	w.Header().Set("Content Type", "application/json")
	w.Write(jsonData)
}