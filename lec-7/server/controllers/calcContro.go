package controllers

import (
	"encoding/json"
	"net/http"

	"../utils"
)

var operatorsArray = make([]string, 5)
var numbersArray = make([]float64, 6)

var errDiv string

type Calc struct {
	Operators []string
	Numbers   []float64
}

func Calculator(w http.ResponseWriter, req *http.Request) {
	var p Calc
	err := json.NewDecoder(req.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	operatorsArray = p.Operators
	numbersArray = p.Numbers

	runClosureMulDiv := utils.CalcMulDiv(operatorsArray, numbersArray, errDiv)
	for i := 0; i <= len(operatorsArray); i++ {

		if errDiv != "" {
			var response = map[string]interface{}{
				"msg": errDiv,
			}
			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				return
			}
			errDiv = ""
			return
		}
		runClosureMulDiv()
	}

	runClosureAddSub := utils.CalcAddSub(operatorsArray, numbersArray, errDiv)
	for i := 0; i <= len(operatorsArray); i++ {
		runClosureAddSub()
	}
	var data = map[string]interface{}{
		"msg": numbersArray[0],
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}

}
