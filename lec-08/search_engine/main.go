package main

import (
	"./models"
	"./routes"
)

func main() {
	models.NewElasticSearchClient()
	models.CheckExistsIndex()

	routes.Init()
}
