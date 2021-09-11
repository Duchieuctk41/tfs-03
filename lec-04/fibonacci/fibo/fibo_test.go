package fibo

import "testing"

func TestFibo(t *testing.T) {
	nInput := int(10)
	expectedOutput := int(55)
	realOutput := Fibo(nInput)

	if expectedOutput != realOutput {
		t.Errorf("Got: %v, but expected: %v", expectedOutput, realOutput)
	}
}
