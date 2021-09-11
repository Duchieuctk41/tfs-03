package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Cataloger interface {
	Create(db *sql.DB) error
}

type Catalog struct {
	ID   int
	Name string
}

func (ca *Catalog) Create(db *sql.DB) error {
	// _, err := db.Exec(`INSERT INTO catalog(id, name) VALUES(?, ?)`, ca.ID, ca.Name)
	insert, err := db.Query("INSERT INTO catalog(id, name) VALUES (?, ?)", ca.ID, ca.Name)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

type Database struct {
	db *sql.DB
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/ecommerce")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	database := &Database{db: db}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/cataloger", database.Welcome).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))

}

func (database *Database) Welcome(w http.ResponseWriter, r *http.Request) {
	catalog := &Catalog{3, "huu"}
	var cataloger Cataloger
	cataloger = catalog
	cataloger.Create(database.db)
	fmt.Fprintf(w, "hello world")
}
