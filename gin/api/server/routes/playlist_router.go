package routes

import (
	"Services"

	"github.com/gin-gonic/gin"
)

func CreatePlaylistRoutes(parentRouter *gin.RouterGroup) {
    //router := PlaylistRouter{}
    playlistRouterGroup := parentRouter.Group("playlist")
    playlistRouterGroup.GET("/all", getPlaylists)
    //playlistRouterGroup.GET()
}

func getPlaylists(c *gin.Context) {
    playlistService := Services.NewPlaylistService()
    data, err := playlistService.GetPlaylists()

    SendResponse(c, 200, data, err)
}

