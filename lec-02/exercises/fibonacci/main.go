package main

import (
	"fmt"

	"./fibo"
)

func main() {
	result := fibo.CalculateFibo_n(10)
	fmt.Println(result)
}
