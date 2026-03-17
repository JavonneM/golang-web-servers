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
		fmt.Println("Tracer")
		reqId, err := uuid.NewV7()
		if err != nil {
			fmt.Println("tracer: internal server error")
			return
		}
		ctx := context.WithValue(r.Context(), ContextRequestTraceKey, reqId)

		fmt.Println(ctx.Value(ContextRequestTraceKey))
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
