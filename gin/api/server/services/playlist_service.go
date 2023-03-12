package Services
import ST "ServiceTypes"
type PlaylistService interface {
    GetPlaylists() []ST.PlaylistResponse
}

