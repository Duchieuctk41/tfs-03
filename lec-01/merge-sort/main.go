package main

import "fmt"

func main() {
	a := []int{3, 2, 1}
	fmt.Println(merge(a))
}

func merge(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	for i := 0; i < len(arr); i++ {
		if i < mid {
			left[i] = arr[i]
		} else {
			right[i-mid] = arr[i]
		}
	}
	return mergeUnsorted(merge(left), merge(right))
}

func mergeUnsorted(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
