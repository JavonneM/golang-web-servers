package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logger")
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)

		fmt.Println(r.Context().Value(ContextRequestTraceKey))
		next(w, r)
	})
}
