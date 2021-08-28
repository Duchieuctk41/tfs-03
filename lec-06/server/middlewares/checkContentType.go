package middlewares

import (
	"fmt"
	"net/http"
)

const JsonContentType = "application/json"

func ContentTypeChecking(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != JsonContentType {
			fmt.Fprintf(w, "request only allow content type application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}
