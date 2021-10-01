package main

import "fmt"

func main() {
	a := []int{3, 2, 1}
	fmt.Println(Bubble(a))
}

func Bubble(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for y := 0; y < len(a)-1-i; y++ {
			if a[y] > a[y+1] {
				a[y], a[y+1] = a[y+1], a[y]
			}
		}
	}
	return a
}
