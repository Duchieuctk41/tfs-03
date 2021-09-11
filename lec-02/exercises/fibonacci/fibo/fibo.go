package main

func Fibo(a int) int {
	if a == 1 || a == 2 {
		return 1
	}

	first, second := 1, 1
	var result int
	for i := 3; i <= a; i++ {
		result = first + second
		first, second = second, first+second
	}

	return result
}
