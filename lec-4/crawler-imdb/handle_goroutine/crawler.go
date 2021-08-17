package handle_goroutine

import (
	"fmt"

	"github.com/gocolly/colly"
)

func (listMovie *ListMovie) Crawl(m Movie) []Movie {
	c := colly.NewCollector()

	c.OnHTML(".lister-list", func(tf *colly.HTMLElement) {
		tf.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			m.Key = len(listMovie.Movies)
			m.Name = e.ChildText(".titleColumn > a")
			m.Year = e.ChildText(".titleColumn > .secondaryInfo")
			m.Rating = e.ChildText(".titleColumn > strong")
			listMovie.Movies = append(listMovie.Movies, m)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Crawling", r.URL)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")

	return listMovie.Movies
}
