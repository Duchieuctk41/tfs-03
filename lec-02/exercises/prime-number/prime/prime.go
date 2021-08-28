package prime

func FindPrimeNumber(a int) bool {
	if a < 2 {
		return false
	}
	if a == 2 {
		return true
	}
	if a%2 == 0 {
		return false
	}

	for i := 3; i < a-1; i += 2 {
		if a%i == 0 {
			return false
		}
	}
	return true
}
