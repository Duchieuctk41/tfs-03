package sorts

import (
	"testing"
)

var expectedOutput = []int{1, 2, 3, 5, 8, 9}

func TestBubbleSort(t *testing.T) {
	nInput := []int{5, 8, 2, 9, 3, 1}

	realOutput := BubbleSort(nInput)

	if !compare(realOutput, expectedOutput) {
		t.Errorf("Got: %v but expected %v", expectedOutput, realOutput)
	}
}

func TestQuicksort(t *testing.T) {
	nInput := []int{5, 8, 2, 9, 3, 1}

	realOutput := Quicksort(nInput)

	if !compare(realOutput, expectedOutput) {
		t.Errorf("Got: %v but expected %v", expectedOutput, realOutput)
	}
}

func TestMergeSort(t *testing.T) {
	nInput := []int{5, 8, 2, 9, 3, 1}

	realOutput := MergeSort(nInput)

	if !compare(realOutput, expectedOutput) {
		t.Errorf("Got: %v but expected %v", expectedOutput, realOutput)
	}
}

func compare(arr1, arr2 []int) bool {
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
