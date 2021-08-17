package main

import (
	"sync"

	"./database"
	"./handle"
)

func main() {
	database.Connect()
	var wg sync.WaitGroup
	handle.Handle(&wg)
}
