package main

import (
	"fmt"
	// "net/http"
	// "github.com/gin-gonic/gin"
	gateway "gateways"
)

func main() {
	test := gateway.Songs
	fmt.Println(test[0].Id)
	/*
	   r := gin.Default()
	   r.GET("/ping", func(c *gin.Context) {
	       c.JSON(http.StatusOK, gin.H {
	           "message": "pong",
	       })
	   })
	   r.Run()
	*/
	//CreateRoutes();
}
