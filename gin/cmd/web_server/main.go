package main

import (
	//"fmt"
	// "net/http"
	//"github.com/gin-gonic/gin"
    Router "routes"
	//gateway "gateways"
)

func main() {
    Router.BuildAllRouters().Run()
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
