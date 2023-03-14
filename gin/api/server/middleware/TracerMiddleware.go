package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TracerMiddleware() gin.HandlerFunc{
    return func (c *gin.Context) {
        fmt.Printf("Generating Trace ID\n")
        c.Set("request_id", "123456")
        c.Next()
    }
}
