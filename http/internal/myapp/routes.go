package myapp

import (
	"fmt"

	"github.com/javonnem/web_server/http/internal/myapp/handler"
	"github.com/javonnem/web_server/http/pkg/server"
)

type Routes struct {
	path    string
	handler server.Middleware
}

var routes = []Routes{
	{
		path:    "/healthcheck",
		handler: handler.HealthCheck,
	},
}

type MyAppServer struct {
	server.HttpServer
}

func (*MyAppServer) NewServer() (*MyAppServer, error) {
	var internalServer *server.HttpServer
	internalServer, err := internalServer.NewHttpServer()
	if err != nil {
		return nil, err
	}
	return &MyAppServer{
		*internalServer,
	}, nil
}

func (s *MyAppServer) RegisterRoutes() error {
	// register routes here
	for _, route := range routes {
		err := s.CreateRawRoute(route.path, route.handler)
		if err != nil {
			return fmt.Errorf("failed to register route %w", err)
		}
	}
	return nil
}

func (s *MyAppServer) StartServer() error {
	return s.Listen()
}
