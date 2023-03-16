package controllers

import (
	"beego/models"
    "errors"
	//"encoding/json"
    AppErrors "apperrors"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Playlist 
type PlaylistController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Playlists 
// @Success 200 {object} models.PlaylistResponse
// @router /all [get]
func (u *PlaylistController) GetAllPlaylists() {
	rawData, err := models.GetAllPlaylists()
    
    if err != nil {
        CreateErrorResponse(&u.Controller, AppErrors.NewControllerError("test", errors.New("test"),  "PlaylistService", 500)) 
        return;
    }
    var playlist []models.PlaylistResponse = make([]models.PlaylistResponse, len(rawData))
    for index, item := range rawData {

        convertedItem, _ := item.PlaylistFromMusicApi()
        playlist[index] = convertedItem
    }

    CreateSuccessfulResponse(&u.Controller, 200, playlist)
}
