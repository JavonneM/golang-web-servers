package gateways
import (
    "api/server/gateways/types"
)
type MusicApi interface {
    getPlaylists() []MusicApiPlaylistResponse 
}


