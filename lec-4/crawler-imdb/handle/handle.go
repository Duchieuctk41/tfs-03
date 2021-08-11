package handle

import (
	_ "github.com/go-sql-driver/mysql"
)

func addMovieToArray() myInterface {
	items := []Movie{}
	item := Movie{}
	listMovie := ListMovie{items}

	listMovie.CrawlerData(item)
	return &listMovie
}

func Handle() {
	var callAddMovieToArray func() myInterface
			callAddMovieToArray = addMovieToArray
			callConnectAndUpdateDB := callAddMovieToArray()
			callConnectAndUpdateDB.connectAndUpdateDB()
}
