package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/javonnem/web_server/http/internal/myapp"
)

func SetupHttpServer() (*myapp.MyAppServer, error) {
	var s *myapp.MyAppServer
	s, err := s.NewServer()
	if err != nil {
		return s, fmt.Errorf("failed to create server %w", err)
	}

	err = s.RegisterRoutes()
	if err != nil {
		return s, fmt.Errorf("failed to register server routes %w", err)
	}
	return s, nil
}

func main() {
	println("main")
	s, err := SetupHttpServer()
	if err != nil {
		panic(err)
	}

	go func() {
		// Start Server
		err = s.StartServer()
		if errors.Is(err, http.ErrServerClosed) {
			println("server closed")
		} else if err != nil {
			fmt.Printf("err? %s", err.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block for signal
	<-sigChan
	println("recieved signal")
	c, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	println("starting shutdown")
	s.Shutdown(c)
	println("forcing cancel")
	cancelCtx()
	// Shutdown consumers
	// Shutdown Server
	// Shutdown DB conn
	println("exiting due to signal")
}
