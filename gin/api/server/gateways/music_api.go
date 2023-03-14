package gateways

import (
    BaseErrors "apperrors"
	gatewayTypes "gatewaytypes"
)

type MusicApi interface {
	GetPlaylists() ([]gatewayTypes.MusicApiPlaylistResponse, BaseErrors.GatewayErrorInterface) 
    GetAllSongs() ([]gatewayTypes.MusicApiSongResponse, BaseErrors.GatewayErrorInterface)
}
