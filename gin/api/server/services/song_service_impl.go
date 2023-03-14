package Services

import (
    BaseErrors "apperrors"
	"ServiceTypes"
	Gateways "gateways"
)

type SongServiceImpl struct {
	musicApi Gateways.MusicApi
}

func NewSongService() SongService {
	return SongServiceImpl{Gateways.NewMusicApi()}
}

func (songService SongServiceImpl) GetSongs() ([]ServiceTypes.SongResponse, BaseErrors.BaseErrorInterface) {
	rawData, err := songService.musicApi.GetAllSongs()
    /// TODO(JavonneM) Move into error guard
    if err != nil {
        return nil, err
    }
	var songlist []ServiceTypes.SongResponse = make([]ServiceTypes.SongResponse, len(rawData))
	for index, item := range rawData {
	    songlist[index] = SongToDomain(item)
	}
    return songlist, nil
}
