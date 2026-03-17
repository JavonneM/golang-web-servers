package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/javonnem/web_server/http/pkg/server/middleware"
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

type HandlerResponse struct {
	ProtocolResponseCode int
	Data                 any
}

func ApiHandler[T any](h func(*http.Request, T) HandlerResponse, validate bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("API Handler")
		fmt.Println(r.Context().Value(middleware.ContextRequestTraceKey))
		var body T
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		res := h(r, body)
		// prepare response
		resAsBytes, err := json.Marshal(res.Data)
		if err != nil {
			// reply with marshaling error
			http.Error(w, "unable to encode response", http.StatusInternalServerError)
			return
		}
		// reply with normal response
		// TODO(JavonneM): set json header
		w.WriteHeader(res.ProtocolResponseCode)
		w.Write(resAsBytes)
	})
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

func (hs *HttpServer) CreateRoute(
	rootCtx context.Context,
	route string,
	handler http.HandlerFunc,
	middleware ...Middleware,
) error {
	if hs == nil {
		return fmt.Errorf("http server reciever nil")
	}
	if hs.Mux == nil {
		return fmt.Errorf("internal http server nil")
	}
	return hs.createRoute(rootCtx, route, handler, middleware...)
}

func (hs *HttpServer) CreateRouteWithApiHandling(
	rootCtx context.Context,
	route string,
	handler http.HandlerFunc,
	middleware ...Middleware,
) error {
	if hs == nil {
		return fmt.Errorf("http server reciever nil")
	}
	if hs.Mux == nil {
		return fmt.Errorf("internal http server nil")
	}

	return hs.createRoute(rootCtx, route, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}, middleware...)
}

const defaultHttpTimeout = 30 * time.Second

func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h.ServeHTTP)
	}
	return h
}

func (hs *HttpServer) createRoute(rootCtx context.Context, route string, handler http.HandlerFunc, middleware ...Middleware) error {
	hs.Mux.Handle(route, ChainMiddleware(handler, middleware...))
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
