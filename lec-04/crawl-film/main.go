package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

type Film struct {
	Name   string
	Rating string
	Year   int
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/top_film")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	wg := sync.WaitGroup{}
	ch := make(chan Film)
	wg.Add(2)
	go CrawlFilm(&wg, ch)
	go PushToDB(db, &wg, ch)
	wg.Wait()

	fmt.Println("done")
}

func CrawlFilm(wg *sync.WaitGroup, ch chan<- Film) {
	defer wg.Done()
	defer close(ch)

	c := colly.NewCollector()
	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) {
		item := Film{}
		item.Name = e.ChildText(".titleColumn > a")
		str := e.ChildText(".titleColumn > .secondaryInfo")
		item.Year, _ = strconv.Atoi(str[1:5])
		item.Rating = e.ChildText(".imdbRating > strong")
		fmt.Println(item.Year)
		ch <- item
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visting ", r.URL)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")

	c.OnError(func(_ *colly.Response, err error) {
		log.Println(err)
	})
}

func PushToDB(db *sql.DB, wg *sync.WaitGroup, ch <-chan Film) {
	defer wg.Done()

	for {
		item, ok := <-ch

		if !ok {
			return
		}
		_, err := db.Exec(`INSERT INTO movies(name, rating, year) VALUES(?,?,?)`, &item.Name, &item.Rating, &item.Year)
		if err != nil {
			panic(err)
		}
	}
}
