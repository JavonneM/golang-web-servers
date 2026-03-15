package server

import "net/http"

type Server interface {
	CreateRawRoute() error
	CreateRouteWithDefaultMiddleware() error
	Listen() error
}

type Middleware func (http.HandlerFunc) http.HandlerFunc
