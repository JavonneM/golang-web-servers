package controllers

import (
	"beego/models"
    "errors"
	//"encoding/json"
    AppErrors "apperrors"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Playlist 
type SongController struct {
	beego.Controller
}

// @Title GetAllSongs
// @Description get all songs 
// @Success 200 {object} models.SongResponse
// @router /all [get]
func (u *SongController) GetAllSongs() {
	rawData, err := models.GetAllSongs()
    if err != nil {
        CreateErrorResponse(&u.Controller, AppErrors.NewControllerError("test", errors.New("test"),  "SongController", 500)) 
        return;
    }
    var songlist []models.SongResponse = make([]models.SongResponse, len(rawData))
    for index, item := range rawData {
        convertedItem, _ := item.SongFromMusicApi()
        songlist[index] = convertedItem
    }
    CreateSuccessfulResponse(&u.Controller, 200, songlist)
}
