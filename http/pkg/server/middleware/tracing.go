package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
)

const ContextRequestTraceKey = "request_trace_id"

func Tracer(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId, err := uuid.NewV7()
		if err != nil {
			fmt.Println("tracer: internal server error")
			return
		}
		context.WithValue(r.Context(), ContextRequestTraceKey, reqId)
		next.ServeHTTP(w, r)
	})
}
