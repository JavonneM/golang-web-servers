package gateways

import (
	gatewayTypes "gatewaytypes"
)

type MusicApi interface {
	getPlaylists() []gatewayTypes.MusicApiPlaylistResponse
}
