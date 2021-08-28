package fibo

import "fmt"

func RunFibo_Goroutine(n int, c, quit chan int) {
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}

func Fibo_Goroutine(c, quit chan int) {
	num1, num2 := 0, 1
	for {
		select {
		case c <- num1:
			num1, num2 = num2, num1+num2
		case <-quit:
			return
		}
	}
}
