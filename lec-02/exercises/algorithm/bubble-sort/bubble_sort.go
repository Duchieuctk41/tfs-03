package bubble_sort

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for y := 0; y < len(arr)-1-i; y++ {
			if arr[y] > arr[y+1] {
				arr[y], arr[y+1] = arr[y+1], arr[y]
			}
		}
	}
	return arr
}
