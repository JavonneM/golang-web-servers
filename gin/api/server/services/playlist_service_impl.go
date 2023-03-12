package Services

import (
	"ServiceTypes"
	Gateways "gateways"
)

type PlaylistServiceImpl struct {
	musicApi Gateways.MusicApi
}

func NewPlaylistService() PlaylistService {
	return PlaylistServiceImpl{Gateways.NewMusicApi()}
}

func (playlistService PlaylistServiceImpl) GetPlaylists() []ServiceTypes.PlaylistResponse {
	rawData := playlistService.musicApi.GetPlaylists()
	var playlist []ServiceTypes.PlaylistResponse = make([]ServiceTypes.PlaylistResponse, len(rawData))
	for index, item := range rawData {
	    playlist[index] = PlaylistToDomain(item)
	}
    return playlist
}
