package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type User struct {
	Email    string
	UserName string
	Password string
}

var index = make(map[string]User)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan User)

	// scan csv & map User
	wg.Add(2)
	go scanCSVFile(&wg, ch)
	go mapUser(&wg, ch)
	wg.Wait()

	// search 
	fmt.Println(searchByEmail("a@gmail.com"))
}

func searchByEmail(str string) interface{} {
	if v, ok := index[str]; ok {
		return v
	} else {
		return "not found"
	}
}

func scanCSVFile(wg *sync.WaitGroup, ch chan<- User) {
	defer wg.Done()
	defer close(ch)
	// open file
	file, err := os.Open("map.csv")

	if err != nil {
		log.Fatalf("error when open file csv%s", err)
	}

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}

		item := User{
			Email:    record[0],
			UserName: record[1],
			Password: record[2],
		}
		index[item.Email] = item
	}
}

func mapUser(wg *sync.WaitGroup, ch <-chan User) {
	defer wg.Done()
	for {
		item, ok := <-ch
		if !ok {
			return
		}
		index[item.Email] = item
	}
}
