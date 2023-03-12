package gateways

import (
	gatewayTypes "gatewaytypes"
)

type MusicApi interface {
	GetPlaylists() []gatewayTypes.MusicApiPlaylistResponse
    GetAllSongs() []gatewayTypes.MusicApiSongResponse
}
