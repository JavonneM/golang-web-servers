package dto

import "time"

type PostDataRequest struct {
	Data string `json:"data"`
}

type PostDataResponse struct {
	TimeStamp time.Time `json:"timestamp"`
	PathId string `json:"id"`
	Data string `json:"data"`
}
