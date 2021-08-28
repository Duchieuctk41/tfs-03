package fibo

import "fmt"

func FindFibo(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	num1, num2 := 1, 1
	var resl int
	for i := 3; i <= n; i++ {
		resl = num1 + num2
		num1, num2 = num2, num1+num2
	}
	return resl
}

func FiboNormal(n int) {
	num1, num2 := 1, 0
	for i := 0; i < n; i++ {
		num1, num2 = num2, num1+num2
		fmt.Println(num1)
	}
}
