package main

import (
	"fmt"

	"./bubble_sort"
	"./merge_sort"
	"./quick_sort"
)

func main() {
	arr := []int{9, 10, 2, 5, 0, 4, 7, 8, 1, 6}
	bubble_sort.PrintBubbleSort(arr)
	arr1 := []int{9, 10, 2, 5, 0, 4, 7, 8, 1, 6}
	merge_sort.PrintMergeSort(arr1)
	arr2 := []int{9, 10, 2, 5, 0, 4, 7, 8, 1, 6}
	fmt.Println(quick_sort.Quicksort(arr2))
}
