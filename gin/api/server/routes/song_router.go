package routes

import (
	"Services"

	"github.com/gin-gonic/gin"
)

func CreateSongRoutes(parentRouter *gin.RouterGroup) {
    playlistRouterGroup := parentRouter.Group("songs")
    playlistRouterGroup.GET("/all", SongGetSongs)
}

func SongGetSongs(c *gin.Context) {
    songService := Services.NewSongService()
    data := songService.GetSongs()

    c.JSON(200, gin.H {
        "data": data,
    })
}

