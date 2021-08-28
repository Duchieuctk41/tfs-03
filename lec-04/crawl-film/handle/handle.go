package handle

import (
	"fmt"
	"log"
	"sync"

	"../database"
	"github.com/gocolly/colly"
)

type Movie struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	Year   string
	Rating string
}

func Handle(wg *sync.WaitGroup) {
	data := make(chan Movie)
	wg.Add(2)
	go crawlerMovie(wg, data)
	go createMoive(wg, data)
	wg.Wait()
	log.Println("success crawl")
}

func crawlerMovie(wg *sync.WaitGroup, ch chan Movie) {
	defer wg.Done()
	defer close(ch)

	c := colly.NewCollector()
	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) {
		item := Movie{}
		item.Name = e.ChildText(".titleColumn > a")
		item.Year = e.ChildText(".titleColumn > .secondaryInfo")
		item.Rating = e.ChildText(".imdbRating > strong")
		ch <- item
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("How", r.URL)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")

	c.OnError(func(_ *colly.Response, err error) {
		log.Println(err)
	})
}

func createMoive(wg *sync.WaitGroup, ch chan Movie) {
	defer wg.Done()
	id := 1
	for {
		m, ok := <-ch
		if !ok {
			return
		}
		m.Id = id
		database.DB.Create(&m)
		id++
	}
}
