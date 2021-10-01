package main

import "fmt"

func main() {
	a := []int{3, 2, 1}
	fmt.Println(quick(a, 0, len(a)-1))
}

//partition
func partion(a []int, low, high int) ([]int, int) {
	pivot := a[high]
	i := low
	for j := low; j < high; j++ {
		if a[j] < a[pivot] {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return a, i
}

func quick(a []int, low, high int) []int {
	if low < high {
		a, p := partion(a, low, high)
		a = quick(a, low, p-1)
		a = quick(a, p+1, high)
	}
	return a
}
