package main

import (
    "fmt"
	"github.com/gin-gonic/gin"
    Router "routes"
    Middleware "middleware"
)

func main() {
    engine := gin.Default()

    fmt.Printf("Enabling Tracing middleware\n")
    engine.Use(Middleware.TracerMiddleware())
    fmt.Printf("Enabling Logging middleware\n")
    engine.Use(Middleware.LoggerMiddleware())

    Router.BuildAllRouters(engine)
    engine.Run()
}
