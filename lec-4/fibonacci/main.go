package main

import (
	"./fibo"
)

func main() {
	n := 6

	// f := fibo_Closure()
	// for i := 0; i < n; i++ {
	// 	fmt.Println(f())
	// }

	c := make(chan int, 3)
	quit := make(chan int)
	go fibo.RunFibo_Goroutine(n, c, quit)
	fibo.Fibo_Goroutine(c, quit)
}
