package utils

import (
	"math"
)

type ErrorString struct {
	s string
}

type CatchError interface {
	Error() string
}

func (e *ErrorString) Error() string {
	return e.s
}

// Hàm nhân chia trước
func CalcMulDiv(operatorsArray []string, numbersArray []float64, errDiv string) func() ([]string, []float64) {
	return func() ([]string, []float64) {
		for i, s := range operatorsArray {
			if s == "*" || s == "/" || s == "%" {
				result := 0.0
				switch s {
				case "*":
					result = numbersArray[i] * numbersArray[i+1]
				case "%":
					if math.Mod(numbersArray[i]/numbersArray[i+1], 1.0) == 0.0 {
						result = 0
					} else {
						result = 1
					}
				case "/":
					// Báo lỗi khi chia cho 0
					if numbersArray[i+1] == 0 {
						var err CatchError = &ErrorString{s: "Phép chia cho 0 thì Chịu, không tính được."}
						errDiv = err.Error()
						return nil, nil
					}
					result = numbersArray[i] / numbersArray[i+1]
				default:
					break

				}
				// xóa phần tử trong mảng operator
				if i == len(operatorsArray)-1 {
					operatorsArray = append(operatorsArray[:i])
				} else {
					operatorsArray = append(operatorsArray[:i], operatorsArray[i+1:]...)
				}
				// xóa phần tử trong mảng numbers
				numbersArray[i] = result
				if i+1 == len(numbersArray)-1 {
					numbersArray = append(numbersArray[:i+1])
				} else {
					numbersArray = append(numbersArray[:i+1], numbersArray[i+2:]...)
				}
				return operatorsArray, numbersArray
			}
		}
		return nil, nil
	}
}

// Hàm tính cộng trừ sau
func CalcAddSub(operatorsArray []string, numbersArray []float64, errDiv string) func() ([]string, []float64) {
	return func() ([]string, []float64) {
		for i, s := range operatorsArray {
			if s == "+" || s == "-" {
				result := 0.0
				switch s {
				case "+":
					result = numbersArray[i] + numbersArray[i+1]
				case "-":
					result = numbersArray[i] - numbersArray[i+1]
				default:
					break
				}

				// xóa phần tử trong mảng operator
				if i == len(operatorsArray)-1 {
					operatorsArray = append(operatorsArray[:i])
				} else {
					operatorsArray = append(operatorsArray[:i], operatorsArray[i+1:]...)
				}
				// xóa phần tử trong mảng numbers
				numbersArray[i] = result
				if i+1 == len(numbersArray)-1 {
					numbersArray = append(numbersArray[:i+1])
				} else {
					numbersArray = append(numbersArray[:i+1], numbersArray[i+2:]...)
				}
				return operatorsArray, numbersArray
			}
		}
		return nil, nil
	}
}
