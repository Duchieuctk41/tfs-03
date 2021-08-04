package main

import (
	"./bubble_sort"
	"./merge_sort"
	// "./quick_sort"
)

func main() {
	arr := [] int { 9, 10, 2, 5, 0, 4, 7, 8, 1, 6 }
	bubble_sort.PrintBubbleSort(arr)
	merge_sort.PrintMergeSort(arr)
}

