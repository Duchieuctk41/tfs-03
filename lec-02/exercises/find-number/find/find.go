package find

func FindBiggestNumber(arr [5] float64) float64 {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func FindSmallesttNumber(arr [5] float64) float64 {
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}

func AverageOfArray(arr [5] float64) float64 {
var a float64
	for _, v := range arr {
		a += v
	}
	return a
}