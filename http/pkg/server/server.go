package server

type Server interface {
	CreateRawRoute() error
	CreateRouteWithDefaultMiddleware() error
	Listen() error
}

type Middleware func(any) (any, error)
