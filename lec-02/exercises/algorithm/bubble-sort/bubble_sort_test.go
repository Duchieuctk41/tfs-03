package main

import "testing"

func TestBubbleSort(t *testing.T) {
	aInput := []int{5, 1, 8, 3, 6}       // giá trị đầu vào
	expectedOutput := []int{1, 3, 5, 6, 8} // giá trị mong đợi

	realOutput := BubbleSort(aInput) // giá trị trả về thực tế

	if !equal(expectedOutput, realOutput) {
		t.Errorf("Got: %v,  but expected: %v", expectedOutput, realOutput)
	}
}

func equal(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i, v := range a1 {
		if v != a2[i] {
			return false
		}
	}
	return true
}
