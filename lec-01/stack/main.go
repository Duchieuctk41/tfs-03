package main

import "fmt"

// stack
type Stack struct {
	array []int
}

// push stack
func (s *Stack) Push(i int) {
	s.array = append(s.array, i)
}

// pop stack
func (s *Stack) Pop() int {
	l := len(s.array) - 1
	toRemove := s.array[l]
	s.array = s.array[:l]
	return toRemove
}

func main() {
	stack := Stack{}
	fmt.Println(stack)
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
