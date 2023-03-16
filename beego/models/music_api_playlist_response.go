package models
import AppErrors "apperrors"
type MusicApiPlaylistResponse struct {
	Id    int
	Title string
	Songs []MusicApiSongResponse
}

var playlistResponse = []MusicApiPlaylistResponse{
	{Id: 0, Title: "Default", Songs: []MusicApiSongResponse{
		songs[0],
		songs[1],
		songs[2],
	},
	},
	{Id: 1, Title: "Special", Songs: []MusicApiSongResponse{
		songs[2],
		songs[3],
		songs[2],
	},
	},
}
var _songs = []MusicApiSongResponse{
	{Id: 1, Title: "Song A", Url: "https://bash.com/test.mp4", Description: "this is song A"},
	{Id: 2, Title: "Song B", Url: "https://bash.com/test.mp4", Description: "this is song B"},
	{Id: 3, Title: "Song C", Url: "https://bash.com/test.mp4", Description: "this is song C"},
	{Id: 4, Title: "Song D", Url: "https://bash.com/test.mp4", Description: "this is song D"},
	{Id: 5, Title: "Song E", Url: "https://bash.com/test.mp4", Description: "this is song E"},
}

func GetAllPlaylists() ([]MusicApiPlaylistResponse, AppErrors.ModelErrorInterface) {
    return playlistResponse, nil
}
