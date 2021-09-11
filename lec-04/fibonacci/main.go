package main

import "fibo/fibo"

func main() {
	n := 10

	// f := fibo.FiboClosure()
	// for i := 0; i <= n; i++ {
	// 	fmt.Println(f())
	// }

	ch := make(chan int, 3)
	out := make(chan bool)
	go fibo.RunFibo(n, ch, out)
	fibo.FiboGoroutine(ch, out)
}
