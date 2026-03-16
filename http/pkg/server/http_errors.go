package server

import "errors"

// Http 400
var ErrHttpNotFoundError = errors.New("not found")
var ErrHttpUnprocessableEntityError = errors.New("unable to process request")
var ErrHttpBadRequestError = errors.New("unable to decode request")

// Http 500
var ErrHttpInternalServerError = errors.New("internal server error")

type HttpError struct {
	Error error `json:"error"`
}
