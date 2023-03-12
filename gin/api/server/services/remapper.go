package Services
import (
	"ServiceTypes"
	GatewayTypes "gatewaytypes"
)
// / mapping functions
//type ApiSongType GatewayTypes.MusicApiSongResponse

func SongToDomain(song GatewayTypes.MusicApiSongResponse) ServiceTypes.SongResponse {
	return ServiceTypes.SongResponse{
		Id:          song.Id,
		Description: song.Description,
	}
}

//type ApiPlaylistType GatewayTypes.MusicApiPlaylistResponse

func PlaylistToDomain(playlist GatewayTypes.MusicApiPlaylistResponse) ServiceTypes.PlaylistResponse {
	var songs = make([]ServiceTypes.SongResponse, len(playlist.Songs))
	for index, item := range playlist.Songs {
		songs[index] = SongToDomain(item)
	}

	return ServiceTypes.PlaylistResponse {
        Id: string(playlist.Id),
        Title: playlist.Title,
        Songs: songs,
    }
}


