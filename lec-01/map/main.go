package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"
)

type User struct {
	Email   string
	Name    string
	Address string
}

var user = make(map[string]User)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan User)

	wg.Add(2)
	go scanFile(&wg, ch)
	go mapUser(&wg, ch)
	wg.Wait()

	search("a@gmail.com")

}

func scanFile(wg *sync.WaitGroup, ch chan<- User) {
	defer wg.Done()
	defer close(ch)

	file, err := os.Open("map.csv")
	if err != nil {
		log.Fatalf("%s", err)
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
			Email:   record[0],
			Name:    record[1],
			Address: record[2],
		}
		ch <- item
	}
}

func mapUser(wg *sync.WaitGroup, ch <-chan User) {
	defer wg.Done()
	for {
		item, ok := <-ch
		if !ok {
			return
		}
		user[item.Email] = item
	}
}

func search(email string) interface{} {
	if v, ok := user[email]; ok {
		return v
	} else {
		return "not found"
	}
}
