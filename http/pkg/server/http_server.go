package server

import (
	"context"
	// "encoding/json"
	"fmt"
	"net/http"
	"time"
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

func (hs *HttpServer) CreateRoute(rootCtx context.Context, route string, handler http.HandlerFunc, middleware ...Middleware) error {
	if hs == nil {
		return fmt.Errorf("http server reciever nil")
	}
	if hs.Mux == nil {
		return fmt.Errorf("internal http server nil")
	}
	return hs.createRoute(rootCtx, route, handler, middleware...)
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

const defaultHttpTimeout = 30 * time.Second

func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h.ServeHTTP)
	}
	return h
}

func (hs *HttpServer) createRoute(rootCtx context.Context, route string, handler http.HandlerFunc, middleware ...Middleware) error {
	hs.Mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		/// TODO(JavonneM): Fix middleware handling
		// var err error
		// var res any

		// for _, mware := range middleware {
		// 	res, err = mware(appContext)
		// 	if err != nil {
		// 		break
		// 	}
		// }
		h := ChainMiddleware(handler, middleware...)
		h.ServeHTTP(w, r)
		// execute handler
		// if err != nil {
		// 	fmt.Println(res)
		// 	fmt.Println(err)
		// }
		// payload, err := json.Marshal(&res)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// // TODO(JavonneM): Write status code back, retrieve from error
		// _, err = w.Write(payload)
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
