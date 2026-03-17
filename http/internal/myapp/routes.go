package myapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/javonnem/web_server/http/internal/myapp/dto"
	"github.com/javonnem/web_server/http/internal/myapp/handler"
	"github.com/javonnem/web_server/http/pkg/server"
	"github.com/javonnem/web_server/http/pkg/server/middleware"
)

type Routes struct {
	path       string
	middleware []server.Middleware
	handler    http.HandlerFunc
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

func (s *MyAppServer) registerSystemRoutes(rootCtx context.Context, sh handler.SystemHandler) error {
	err := s.CreateRoute(rootCtx, "/v1/healthcheck", sh.HealthCheck, middleware.Tracer, middleware.Logger)
	if err != nil {
		return fmt.Errorf("failed to register healthcheck route %w", err)
	}

	err = s.CreateRouteWithApiHandling(rootCtx, "/v1/testapihandler", server.ApiHandler[dto.TestApiHandlerRequest](sh.TestApiHandler, false), middleware.Tracer, middleware.Logger)
	if err != nil {
		return fmt.Errorf("failed to register healthcheck route %w", err)
	}

	err = s.CreateRouteWithApiHandling(rootCtx, "/v2/testapihandler/{id}", server.ApiHandler[any](sh.TestApiHandlerWithPathValues, false), middleware.Tracer, middleware.Logger)
	if err != nil {
		return fmt.Errorf("failed to register healthcheck route %w", err)
	}

	return nil
}

func (s *MyAppServer) registerAppRoutes(rootCtx context.Context, sh handler.SystemHandler) error {
	return nil
	//err := s.CreateRawRoute(rootCtx, "/healthcheck", middleware.Logger)
	//if err != nil {
	//	return fmt.Errorf("failed to register healthcheck route %w", err)
	//}

	// return nil
	//	path:       "/account/:id",
	//	middleware: []server.Middleware{middleware.Logger, middleware.Tracer},
	//	handler: func(w http.ResponseWriter, r *http.Request) (any, error) {
	//		return handler.PostDataWithPathParameters(ctx, nil)
	//	},
}

func (s *MyAppServer) RegisterRoutes(rootCtx context.Context, sh handler.SystemHandler) error {
	// Bind system endpoints
	err := s.registerSystemRoutes(rootCtx, sh)
	if err != nil {
		return fmt.Errorf("failed to register healthcheck route %w", err)
	}
	return nil
}

func (s *MyAppServer) StartServer() error {
	return s.Listen()
}
