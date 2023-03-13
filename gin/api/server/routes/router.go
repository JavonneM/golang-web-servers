package routes

import (
    "github.com/gin-gonic/gin"
)

func BuildAllRouters() *gin.Engine {
    engine := gin.Default()
    buildV1Routers(engine)
    return engine;
}

func buildV1Routers(engine *gin.Engine) {
    versionRouterGroup := engine.Group("v1")
    CreatePlaylistRoutes(versionRouterGroup)
    CreateSongRoutes(versionRouterGroup)
    //SongRouter.CreateRoutes(versionRouterGroup)
}
