package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Movie struct {
	Name string `json:"name"`
	Year string `json:"year"`
	Rating string `json:"rating"`
}

type ListMovie []Movie
var tmpListMovie = ListMovie{}

func CrawlerFromIMDB() interface{} {
	c := colly.NewCollector()
	c.OnHTML(".lister-list", func(tf *colly.HTMLElement) {
		tf.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			tmpMovie := Movie{}
			tmpMovie.Name = e.ChildText(".titleColumn > a")
			tmpMovie.Year = e.ChildText(".titleColumn > .secondaryInfo")
			tmpMovie.Rating = e.ChildText(".imdbRating > strong")
			tmpListMovie = append(tmpListMovie, tmpMovie)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("How", r.URL)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
	return tmpListMovie
}