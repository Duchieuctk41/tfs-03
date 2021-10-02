package main

import "fmt"

func main() {
	a := []int{3, 2, 1}
	fmt.Println(quickSort(a, 0, len(a)-1))
}

func partition(a []int, low, high int) ([]int, int) {
	pivot := a[high]
	i := low
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return a, i
}

func quickSort(a []int, low, high int) []int {
	if low < high {
		a, p := partition(a, low, high)
		a = quickSort(a, low, p-1)
		a = quickSort(a, p+1, high)
	}
	return a
}
