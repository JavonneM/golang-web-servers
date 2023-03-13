package Services
import ST "ServiceTypes"
type SongService interface {
    GetSongs() []ST.SongResponse
}

