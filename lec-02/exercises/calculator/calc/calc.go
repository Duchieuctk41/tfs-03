package calc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var (
	operator = "type_of_operator"
	num1 = 1.0
	num2 = 1.0
)

func Calculator(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	fmt.Println(params)
	if operatorType, ok := params["operator"]; ok {
		operator = operatorType[0]
	}
	if number1, ok := params["num1"]; ok {
		num1, _ = strconv.ParseFloat(number1[0], 64)
	}
	if number2, ok := params["num2"]; ok {
		num2, _ = strconv.ParseFloat(number2[0], 64)
	}
	result := Calc(operator, num1, num2)
	var data = map[string]interface{}{
		"msg": result,
	}
	json.NewEncoder(w).Encode(data)
}


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