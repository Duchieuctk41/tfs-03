package main

import (
	"fmt"
	"./find"
)

func main() {
	num := [...] float64{2,9,1,5,4}
	max := find.FindBiggestNumber(num)
	fmt.Println("max: ", max)

	min := find.FindSmallesttNumber(num)
	fmt.Println("min: ", min)

	average := find.AverageOfArray(num)
	fmt.Println("average: ", average)
}