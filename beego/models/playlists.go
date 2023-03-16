package models
import AppErrors "apperrors"
type PlaylistResponse struct {
    Id string
    Title string
    Songs []SongResponse
}

func (playlist MusicApiPlaylistResponse) PlaylistFromMusicApi() (PlaylistResponse, AppErrors.ModelErrorInterface) {
    var songs = make([]SongResponse, len(playlist.Songs))
    for index, item := range playlist.Songs {
        song, _ := item.SongFromMusicApi()
        songs[index] = song
    }

	return PlaylistResponse {
        Id: string(playlist.Id),
        Title: playlist.Title,
        Songs: songs,
    }, nil

}
