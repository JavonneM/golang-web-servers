package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/go-playground/sensitive"
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

const (
	ModeMigration  = "migration"
	ModeHttpServer = "server"
)

type ConfiguredFlags struct {
	Mode RunMode
}

type RunMode struct {
	mode string
}

func (r *RunMode) String() string {
	return r.mode
}

func (r *RunMode) Set(s string) error {
	switch s {
	case ModeMigration, ModeHttpServer:
		r.mode = s
		return nil
	}
	return errors.New("unsupported run mode")
}

func setupFlags() (ConfiguredFlags, error) {
	configuredFlags := ConfiguredFlags{}
	fs := flag.NewFlagSet("eg", flag.ExitOnError)
	fs.Var(&configuredFlags.Mode, "mode", "mode description")
	err := fs.Parse(os.Args[1:])
	return configuredFlags, err
}

type EnvVars struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername sensitive.String
	DatabasePassword sensitive.String
}

func configureEnvs() EnvVars {
	v := EnvVars{}
	return v
}

func main() {
	flags, err := setupFlags()
	_ = configureEnvs()
	if err != nil {
		os.Exit(1)
	}
	if flags.Mode.String() == ModeMigration {
		println("Migration")
	} else if flags.Mode.String() == ModeHttpServer {
		println("Server mode")
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
}
