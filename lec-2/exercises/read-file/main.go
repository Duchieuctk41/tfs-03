package main

import (
	"fmt"

	// import 
	"./read"
)

func main() {
	var fileName string = "./text/text.txt" 
	read.ReadFile(fileName)

	fmt.Println("\n=================================")
	var str string = "say hello\n"
	read.WrFile(fileName, str)
	read.ReadFile(fileName)


}