package calculator

func Calc(operator string, a float64, b float64) float64 {
	switch operator {
	case "mul":
		return a * b
	case "did":
		if b == 0 {
			return 0
		}
		return a / b
	case "sum":
		return a + b
	case "sub":
		return a - b
	default:
		return 0
	}
}