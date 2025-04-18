package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AppServer interface {
	RegisterRoutes() error
	StartServer() error
}

type HttpServer struct {
	Port                   string
	Mux                    *http.ServeMux
	Server                 *http.Server
	DefaultMiddlewareStack []Middleware
}

func (hs *HttpServer) NewHttpServer() (*HttpServer, error) {
	if hs != nil {
		return nil, fmt.Errorf("unable to start a new server using current instance")
	}
	mux := http.NewServeMux()
	return &HttpServer{
		Port: ":8080",
		Mux:  mux,
	}, nil
}

func (hs *HttpServer) CreateRawRoute(route string, handler Middleware) error {
	if hs == nil {
		return fmt.Errorf("http server reciever nil")
	}
	if hs.Mux == nil {
		return fmt.Errorf("internal http server nil")
	}
	return hs.createRoute(route, []Middleware{handler})
}

func (hs *HttpServer) CreateRouteWithDefaultMiddleware() error {
	if hs == nil {
		return fmt.Errorf("http server reciever nil")
	}
	if hs.Mux == nil {
		return fmt.Errorf("internal http server nil")
	}
	return nil
}

func (hs *HttpServer) createRoute(route string, middleware []Middleware) error {
	hs.Mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		var err error
		var res any
		for _, mware := range middleware {
			res, err = mware(nil)
			if err != nil {
				break
			}
		}

		// execute handler
		if err != nil {
			println(res)
		}
		payload, err := json.Marshal(&res)

		_, err = w.Write(payload)
	})
	return nil
}

func (hs *HttpServer) Listen() error {
	if hs.Server != nil {
		return fmt.Errorf("server already started")
	}
	hs.Server = &http.Server{
		Addr:    hs.Port,
		Handler: hs.Mux,
	}
	return hs.Server.ListenAndServe()
}

func (hs *HttpServer) Shutdown(c context.Context) error {
	return hs.Server.Shutdown(c)
}
