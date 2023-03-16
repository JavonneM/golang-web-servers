package models

import (
    AppErrors "apperrors"
)
type SongResponse struct {
    Id int
    Description string
}

func (song MusicApiSongResponse) SongFromMusicApi() (SongResponse, AppErrors.ModelErrorInterface) {
    return SongResponse {
        Id: song.Id,
        Description: song.Description,
    }, nil
}
