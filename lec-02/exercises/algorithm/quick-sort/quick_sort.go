package main

import (
	"fmt"
	"math/rand"
)

func Quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		// fmt.Println(a[i]," ", a[right]-1, " ", i, " ", left)
		if a[i] < a[right]-1 { //a[right] - 1 == a[6] - 1
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	// recursive
	Quicksort(a[:left])
	Quicksort(a[left+1:])

	return a
}

func main() {
	a := []int{5, 1, 8, 3, 6}
	Quicksort(a)
	fmt.Println(a)
}
