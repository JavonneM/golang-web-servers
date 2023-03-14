package Services
import (
    BaseErrors "apperrors"
    ST "ServiceTypes"
)
type PlaylistService interface {
    GetPlaylists() ([]ST.PlaylistResponse, BaseErrors.ServiceErrorInterface)
}

