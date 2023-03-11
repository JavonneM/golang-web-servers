package gateways

import gatewayTypes "gatewaytypes"

var PlaylistResponse = []gatewayTypes.MusicApiPlaylistResponse{
	{Id: 0, Title: "Default", Songs: []gatewayTypes.MusicApiSongResponse{
		Songs[0],
		Songs[1],
		Songs[2],
	},
	},
	{Id: 1, Title: "Special", Songs: []gatewayTypes.MusicApiSongResponse{
		Songs[2],
		Songs[3],
		Songs[2],
	},
	},
}
var Songs = []gatewayTypes.MusicApiSongResponse{
	{Id: 1, Title: "Song A", Url: "https://bash.com/test.mp4", Description: "this is song A"},
	{Id: 2, Title: "Song B", Url: "https://bash.com/test.mp4", Description: "this is song B"},
	{Id: 3, Title: "Song C", Url: "https://bash.com/test.mp4", Description: "this is song C"},
	{Id: 4, Title: "Song D", Url: "https://bash.com/test.mp4", Description: "this is song D"},
	{Id: 5, Title: "Song E", Url: "https://bash.com/test.mp4", Description: "this is song E"},
}

// func (musicApi MusicApi) GetPlaylists() []gatewayTypes.MusicApiPlaylistResponse {
// 	return PlaylistResponse
// }
