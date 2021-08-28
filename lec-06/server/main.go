package main

import (
	"./database"
	"./routes"
)

func main() {
	database.Connect()
	routes.Init()
}
