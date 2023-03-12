package gateways

import gatewayTypes "gatewaytypes"

var playlistResponse = []gatewayTypes.MusicApiPlaylistResponse{
	{Id: 0, Title: "Default", Songs: []gatewayTypes.MusicApiSongResponse{
		songs[0],
		songs[1],
		songs[2],
	},
	},
	{Id: 1, Title: "Special", Songs: []gatewayTypes.MusicApiSongResponse{
		songs[2],
		songs[3],
		songs[2],
	},
	},
}
var songs = []gatewayTypes.MusicApiSongResponse{
	{Id: 1, Title: "Song A", Url: "https://bash.com/test.mp4", Description: "this is song A"},
	{Id: 2, Title: "Song B", Url: "https://bash.com/test.mp4", Description: "this is song B"},
	{Id: 3, Title: "Song C", Url: "https://bash.com/test.mp4", Description: "this is song C"},
	{Id: 4, Title: "Song D", Url: "https://bash.com/test.mp4", Description: "this is song D"},
	{Id: 5, Title: "Song E", Url: "https://bash.com/test.mp4", Description: "this is song E"},
}
type MusicApiImpl struct {
    allSongs []gatewayTypes.MusicApiSongResponse
    allPlaylists []gatewayTypes.MusicApiPlaylistResponse 
}

func NewMusicApi() MusicApi {
    return MusicApiImpl {
        allSongs: songs,
        allPlaylists: playlistResponse,
    }
}



func (musicApi MusicApiImpl) GetPlaylists() []gatewayTypes.MusicApiPlaylistResponse {
 	return musicApi.allPlaylists
}

func (musicApi MusicApiImpl) GetAllSongs() []gatewayTypes.MusicApiSongResponse {
 	return musicApi.allSongs
}
