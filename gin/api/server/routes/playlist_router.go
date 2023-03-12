package routes

import (
	"Services"

	"github.com/gin-gonic/gin"
)

func CreateRoutes(parentRouter *gin.RouterGroup) {
    //router := PlaylistRouter{}
    playlistRouterGroup := parentRouter.Group("playlist")
    playlistRouterGroup.GET("/all", getPlaylists)
    //playlistRouterGroup.GET()
}

func getPlaylists(c *gin.Context) {
    playlistService := Services.NewPlaylistService()
    data := playlistService.GetPlaylists()

    c.JSON(200, gin.H {
        "data": data,
    })
}

