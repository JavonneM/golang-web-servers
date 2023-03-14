package Services

import (
    "errors"
    BaseErrors "apperrors"
	"ServiceTypes"
	Gateways "gateways"
)

type PlaylistServiceImpl struct {
	musicApi Gateways.MusicApi
}

func NewPlaylistService() PlaylistService {
	return PlaylistServiceImpl{Gateways.NewMusicApi()}
}

func (playlistService PlaylistServiceImpl) GetPlaylists() ([]ServiceTypes.PlaylistResponse, BaseErrors.ServiceErrorInterface) {
	rawData, err := playlistService.musicApi.GetPlaylists()
    if err != nil {
        return nil, BaseErrors.NewServiceError("test", errors.New("test"),  "PlaylistService", 500) 
    }
	var playlist []ServiceTypes.PlaylistResponse = make([]ServiceTypes.PlaylistResponse, len(rawData))
	for index, item := range rawData {
	    playlist[index] = PlaylistToDomain(item)
	}
    return playlist, nil
}
