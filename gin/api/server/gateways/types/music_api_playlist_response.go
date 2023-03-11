package gatewaytypes

type MusicApiPlaylistResponse struct {
	Id    int
	Title string
	Songs []MusicApiSongResponse
}
