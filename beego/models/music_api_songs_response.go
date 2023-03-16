package models
import(
    AppErrors "apperrors"
    "errors"
)

type MusicApiSongResponse struct {
	Id          int
	Title       string
	Url         string
	Description string
}

var songs = []MusicApiSongResponse{
	{Id: 1, Title: "Song A", Url: "https://bash.com/test.mp4", Description: "this is song A"},
	{Id: 2, Title: "Song B", Url: "https://bash.com/test.mp4", Description: "this is song B"},
	{Id: 3, Title: "Song C", Url: "https://bash.com/test.mp4", Description: "this is song C"},
	{Id: 4, Title: "Song D", Url: "https://bash.com/test.mp4", Description: "this is song D"},
	{Id: 5, Title: "Song E", Url: "https://bash.com/test.mp4", Description: "this is song E"},
}

func GetAllSongs() ([]MusicApiSongResponse, AppErrors.ModelErrorInterface) {
    return nil, AppErrors.NewModelError(500, "Unable to retrieve songs", errors.New("Raw error obj"), "Domain") 
}
