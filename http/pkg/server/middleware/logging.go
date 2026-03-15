package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next(w, r)
	})
}
