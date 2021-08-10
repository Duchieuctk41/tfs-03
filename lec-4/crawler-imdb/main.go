package main

import (
	// import local package
	"./crawl"
)

func main() {
	crawl.CrawlerFromIMDB()
	crawl.Connect()
}