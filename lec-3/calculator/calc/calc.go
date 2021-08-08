package calc

import (
	"encoding/json"
	"math"
	"net/http"
)

var operatorsArray = make([]string, 5)
var numbersArray = make([]float64, 6)

type Calc struct {
	Operators []string
	Numbers   []float64
}

type ErrorString struct {
	s string
}

type CatchError interface {
	Error() string
}

func (e *ErrorString) Error() string {
	return e.s
}

var errDiv string

func calcMultiAndDivide() func() ([]string, []float64) {

	return func() ([]string, []float64) {
		for i, s := range operatorsArray {
			if s == "*" || s == "/" || s == "%" {
				result := 0.0
				if s == "*" {
					result = numbersArray[i] * numbersArray[i+1]
				}
				if s == "%" {
					if math.Mod(numbersArray[i]/numbersArray[i+1], 1.0) == 0.0 {
						result = 0
					} else {
						result = 1
					}
				}
				if s == "/" {
					if numbersArray[i+1] == 0 {
						var err CatchError = &ErrorString{s: "Phép chia cho 0 thì Chịu, không tính được."}
						errDiv = err.Error()
						return nil, nil
					}
					result = numbersArray[i] / numbersArray[i+1]
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

func calcAddAndSub() func() ([]string, []float64) {
	return func() ([]string, []float64) {
		for i, s := range operatorsArray {
			if s == "+" || s == "-" {
				result := 0.0
				if s == "+" {
					result = numbersArray[i] + numbersArray[i+1]
				}
				if s == "-" {
					result = numbersArray[i] - numbersArray[i+1]
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

func AddCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Set("Access-Control-Allow-Origin", "*")
	headers.Set("Access-Control-Allow-Headers", "Content-Type")
}

func Calculator(w http.ResponseWriter, req *http.Request) {
	AddCorsHeader(w)
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var p Calc
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(req.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	operatorsArray = p.Operators
	numbersArray = p.Numbers

	runClosureMulDiv := calcMultiAndDivide()
	for i := 0; i <= len(operatorsArray); i++ {

		if errDiv != "" {
			var response = map[string]interface{}{
				"msg": errDiv,
			}
			json.NewEncoder(w).Encode(response)
			errDiv = ""
			return
		}
		runClosureMulDiv()
	}

	runClosureAddSub := calcAddAndSub()
	for i := 0; i <= len(operatorsArray); i++ {
		runClosureAddSub()
	}
	var data = map[string]interface{}{
		"msg": numbersArray[0],
	}
	json.NewEncoder(w).Encode(data)

}
