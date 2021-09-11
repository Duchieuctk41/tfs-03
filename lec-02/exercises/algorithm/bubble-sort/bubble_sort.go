package main

func BubbleSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for y := 0; y < len(a)-1-i; y++ {
			if a[y] > a[y+1] {
				a[y], a[y+1] = a[y+1], a[y]
			}
		}
	}
	return a
}
