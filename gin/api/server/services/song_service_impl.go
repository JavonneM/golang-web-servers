package Services

import (
	"ServiceTypes"
	Gateways "gateways"
)

type SongServiceImpl struct {
	musicApi Gateways.MusicApi
}

func NewSongService() SongService {
	return SongServiceImpl{Gateways.NewMusicApi()}
}

func (songService SongServiceImpl) GetSongs() []ServiceTypes.SongResponse {
	rawData := songService.musicApi.GetAllSongs()
	var songlist []ServiceTypes.SongResponse = make([]ServiceTypes.SongResponse, len(rawData))
	for index, item := range rawData {
	    songlist[index] = SongToDomain(item)
	}
    return songlist
}
