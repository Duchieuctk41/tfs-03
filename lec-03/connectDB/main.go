package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Movie struct {
	Name string `json:"name"`
	Year string `json:"year"`
	Rating string `json:"rating"`
}

type ListMovie []Movie

func main() {
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "root:thuyduong1112001@@tcp(127.0.0.1:3306)/movies")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Prepare("INSERT INTO movie(name, year, rating) VALUES(?, ?, ?)")
	//result, err := db.Query("SELECT name FROM users")
	if err != nil {
		panic(err.Error())
	}
	listMovie := ListMovie{
		{"hieu", "1980", "5"},
		{"hoc", "1999", "7"},
	}

	for _,v := range(listMovie) {
		res, err := insert.Exec(v.Name, v.Year, v.Rating)

		if err != nil {
			panic(err.Error())
		}

		id, err := res.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Insert name", id)

	}


	//for result.Next() {
	//	var tag Tag
	//	err = result.Scan(&tag.Name)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	log.Println(tag.Name)
	//}
}
