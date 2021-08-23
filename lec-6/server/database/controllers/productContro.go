package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../database"
	"../models"
	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Fprintf(w, "error when body parse %v", product)
		return
	}

	database.DB.Create(&product)

	fmt.Fprintf(w, "create product: %v", product)
}

func AllsProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	database.DB.Find(&products)

	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var product models.Product

	database.DB.Where("id = ?", id).First(&product)
	if product.Id == 0 {
		fmt.Fprintf(w, "not found product_id: %v", id)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var product models.Product

	database.DB.Where("id = ?", id).First(&product)
	if product.Id == 0 {
		fmt.Fprintf(w, "not found product_id: %v", id)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Fprintf(w, "error when parse body")
		return
	}

	fmt.Println(product)

	database.DB.Model(&product).Updates(product)

	fmt.Fprintf(w, "updated product_id: %v", product.Id)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	fmt.Fprintf(w, "deleted product_id:")
}
