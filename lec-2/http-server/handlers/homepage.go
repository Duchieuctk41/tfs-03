package handlers

import (
	"encoding/json"
	"net/http"
)

func HomePage(w http.ResponseWriter, req *http.Request) {
	var data = map[string]interface{}{
		"msg": "Hello world",
	}
	json.NewEncoder(w).Encode(data)
}
