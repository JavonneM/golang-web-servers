package gateways

import gatewayTypes "types"
var PlaylistResponse = []MusicApiPlaylistResponse{
    {id: 0, title: "Default", songs: []MusicApiSongResponse{
        Songs[0],
        Songs[1],
        Songs[2],
        },
    },
    {id: 1, title: "Special", songs: []MusicApiSongResponse{
        Songs[2],
        Songs[3],
        Songs[2],
        },
    },

}

func (musicApi MusicApi) GetPlaylists() []MusicApiPlaylistResponse {
    return PlaylistResponse
}


