package main

import (
	"./database"
	"./routes"
)

func main() {
	database.ConnectDB()
	routes.Init()
}
