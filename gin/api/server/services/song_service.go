package Services
import (
    BaseErrors "apperrors"
    ST "ServiceTypes"
)
type SongService interface {
    GetSongs() ([]ST.SongResponse, BaseErrors.BaseErrorInterface)
}

