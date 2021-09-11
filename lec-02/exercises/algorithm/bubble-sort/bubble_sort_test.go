package bubble_sort

import "testing"

func TestBubbleSort(t *testing.T) {
	arrInput := []int{5, 1, 8, 3, 6}       // giá trị đầu vào
	expectedOutput := []int{1, 3, 5, 6, 8} // giá trị mong đợi

	realOutput := BubbleSort(arrInput) // giá trị trả về thực tế

	if !equal(expectedOutput, realOutput) {
		t.Errorf("Got: %v,  but expected: %v", expectedOutput, realOutput)
	}
}

func equal(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}
