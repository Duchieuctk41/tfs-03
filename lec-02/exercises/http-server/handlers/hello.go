package handlers

import (
	"fmt"
	"net/http"
)

const (
	nameKey = "your_name"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	if yourName, ok := params["nameKey"]; ok {
		fmt.Fprintf(w, "Hi %s", yourName[0])
	} else {
		fmt.Fprintln(w, `Hi guys. I don't know your name because you don't enter the your_name query param`)
	}
}
