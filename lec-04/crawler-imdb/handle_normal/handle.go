package handle_normal

import (
	_ "github.com/go-sql-driver/mysql"
)

func addMovieToArray() Databaser {
	movies := []Movie{}
	item := Movie{}
	listMovie := ListMovie{movies}

	listMovie.CrawlerData(item)
	return &listMovie
}

func Handle() {
	var callAddMovieToArray func() Databaser
	callAddMovieToArray = addMovieToArray

	callConnectAndUpdateDB := callAddMovieToArray()
	callConnectAndUpdateDB.connectAndUpdateDB()
}
