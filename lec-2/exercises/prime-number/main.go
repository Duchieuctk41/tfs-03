package main

import (
	"fmt"

	"./prime"
)

func main() {
	var num int = 5
	result := prime.FindPrimeNumber(num)
	if result {
		fmt.Println(num, " la so nguyen to")
	} else {
		fmt.Println(num, " khong la so nguyen to")
	}
}
