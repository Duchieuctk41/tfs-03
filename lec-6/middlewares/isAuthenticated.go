package middlewares

import (
	"fmt"
	"net/http"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)

		_, err := r.Cookie("logged_in")

		if err != nil {
			fmt.Fprintf(w, "unauthenticated!")
			return
		}
	})
}
