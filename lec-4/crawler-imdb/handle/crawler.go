package handle

import (
	"fmt"
	"github.com/gocolly/colly"
)

func (listMovie *ListMovie) CrawlerData(item Movie) []Movie {
	c := colly.NewCollector()
	c.OnHTML(".lister-list", func(tf *colly.HTMLElement) {
		tf.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			item.Key = len(listMovie.Items)
			item.Name = e.ChildText(".titleColumn > a")
			item.Year = e.ChildText(".titleColumn > .secondaryInfo")
			item.Rating = e.ChildText(".imdbRating > strong")
			listMovie.Items = append(listMovie.Items, item)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("How", r.URL)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")

	return listMovie.Items
}
