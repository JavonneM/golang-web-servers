package routes

import (
    "github.com/gin-gonic/gin"
)
func CreateRoutes() gin.Engine {
    router := gin.Default()

    //playlistRouter := router.Group("v1")


    return *router;
}
