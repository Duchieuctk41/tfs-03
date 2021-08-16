package handle

import (
	_ "github.com/go-sql-driver/mysql"
)

func addMovieToArray() Databaser {
	items := []Movie{}
	item := Movie{}
	listMovie := ListMovie{items}

	listMovie.CrawlerData(item)
	return &listMovie
}

func Handle() {
	var callAddMovieToArray func() Databaser
	callAddMovieToArray = addMovieToArray
	
	callConnectAndUpdateDB := callAddMovieToArray()
	callConnectAndUpdateDB.connectAndUpdateDB()
}
