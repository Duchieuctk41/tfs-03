package fibo

import "fmt"

func Fibo(a int) int {
	if a == 1 || a == 2 {
		return 1
	}

	first, second := 1, 1
	result := 0

	for i := 3; i <= a; i++ {
		result = first + second
		first, second = second, first+second
	}
	return result
}

func FiboClosure() func() int {
	n := 0
	first, second := 0, 1

	return func() int {
		var result int
		switch {
		case n == 0:
			n++
			result = 0
		case n == 1:
			n++
			result = 1
		default:
			result = first + second
			first, second = second, first+second
		}
		return result
	}
}

func RunFibo(a int, ch <-chan int, out chan<- bool) {
	for i := 0; i <= a; i++ {
		fmt.Println(<-ch)
	}
	out <- false
}

func FiboGoroutine(ch chan<- int, out <-chan bool) int {
	first, second := 0, 1
	for {
		select {
		case ch <- first:
			first, second = second, first+second
		case <-out:
			return 1
		}
	}
}
