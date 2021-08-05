package handlers

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}
