package bubble_sort

import "fmt"

func BubbleSort(arr [] int) [] int{
	for i := 0; i < len(arr) - 1; i++ {
		for y := 0; y < len(arr) -1-i; y++ {
			if arr[y] > arr[y+1] {
				arr[y], arr[y+1] = arr[y+1], arr[y]
			}
		}
	}
	return arr 
}

func PrintBubbleSort(arr[] int) {
	fmt.Println("=========== Before buble sorting =============")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println("\n=========== After bubble sorting =============")
	bub := BubbleSort(arr)
	for _, v := range bub {
		fmt.Print(v, " ")
	}
}