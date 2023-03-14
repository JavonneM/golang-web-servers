package middleware
import (
    "fmt"
    "github.com/gin-gonic/gin"
)
func LoggerMiddleware() gin.HandlerFunc {
    return func (c * gin.Context) {
        fmt.Printf("RequestLogger: %s\n", c.Request.URL.String())
        c.Next()
        fmt.Printf("RequestLogger: After\n") 
    }
}
