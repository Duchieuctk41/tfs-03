package fibo

func Fibo_Closure() func() int {
	n := 0
	num1, num2 := 0, 1

	return func() int {
		var resl int
		switch {
		case n == 0:
			n++
			resl = 0
		case n == 1:
			n++
			resl = 1
		default:
			resl = num1 + num2
			num1, num2 = num2, num1+num2

		}
		return resl
	}
}
