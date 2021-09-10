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

func (y *Catalog) Create(db *sql.DB) error {
	// _, err := db.Exec(`INSERT INTO catalog(id, name) VALUES(id = ?, name = ?)`, &y.ID, &y.Name)
	insert, err := db.Query("INSERT INTO catalog(id, name) VALUES (?, ?)", y.ID, y.Name)
	if err != nil {
		return err
	}
	fmt.Println(insert)
	defer insert.Close()
	return nil
}

type Users struct {
	db *sql.DB
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/ecommerce")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	users := &Users{db: db}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/cataloger", users.Welcome).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))

}

func (users *Users) Welcome(w http.ResponseWriter, r *http.Request) {
	rou := &Catalog{3, "huu"}
	var ca Cataloger
	ca = rou
	ca.Create(users.db)
	fmt.Fprintf(w, "hello world")
}
