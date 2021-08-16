package handle

import (
	"database/sql"
	"fmt"
)

func (listMovie *ListMovie) connectAndUpdateDB() {
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "root:thuyduong1112001@@tcp(127.0.0.1:3306)/movies")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO movie(name, year, rating) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	interfaces := make([]Movier, len(listMovie.Items))

	for k, v := range listMovie.Items {
		interfaces[k] = interface{}(v).(Movier)
	}

	for _, v := range listMovie.Items {
		res, err := insert.Exec(v.Name, v.Year, v.Rating)
		if err != nil {
			panic(err.Error())
		}

		id, err := res.LastInsertId()
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Insert a movie", id)
	}
	// PrintInterfaceSlice(interfaces)
}

// Println data
// func PrintInterfaceSlice(list []Movier) {
// 	for _, v := range list {
// 		fmt.Println(v.key(), v.name(), v.year(), v.rating())
// 	}
// }
