package routes

import (
    ServerErrors "apperrors"
    "github.com/gin-gonic/gin"
)

func BuildAllRouters(engine *gin.Engine) {
    buildV1Routers(engine)
}

func buildV1Routers(engine *gin.Engine) {
    versionRouterGroup := engine.Group("v1")
    CreatePlaylistRoutes(versionRouterGroup)
    CreateSongRoutes(versionRouterGroup)
    //SongRouter.CreateRoutes(versionRouterGroup)
}
func CreateSuccessfulResponse(c *gin.Context, code int, data interface{}) {
    id, _ := c.Get("request_id")
    c.JSON(code, gin.H {
        "success" : true,
        "data": data,
        "message": nil,
        "request_id": id,
    })
}

func CreateErrorResponse(c *gin.Context, err ServerErrors.BaseErrorInterface) {
    id, _ := c.Get("request_id")
    c.Error(err)
    c.JSON(err.ErrorStatusCode(), gin.H {
        "success": false,
        "data" : nil,
        "message": err.ErrorMessage(),
        "request_id": id,
    })
}

func SendResponse(c *gin.Context, successCode int, data interface{}, err ServerErrors.BaseErrorInterface) {
    
    if err != nil {
        CreateErrorResponse(c, err)
    } else {
        CreateSuccessfulResponse(c, successCode, data)
    }
}
