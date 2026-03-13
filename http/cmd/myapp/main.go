package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/sensitive"

	"github.com/javonnem/web_server/http/internal/myapp"
	"github.com/javonnem/web_server/http/pkg/database"
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
	DatabaseName     string
	DatabasePort     string
	DatabaseUsername sensitive.String
	DatabasePassword sensitive.String
}

var ErrEnvMissingVariable = errors.New("missing env var")

func configureEnvs() (EnvVars, error) {
	v := EnvVars{}
	var exists bool
	v.DatabaseHost, exists = os.LookupEnv("DATABASE_HOST")
	if exists && v.DatabaseHost != "" {
		return v, fmt.Errorf("database host missing %w", ErrEnvMissingVariable)
	}
	v.DatabaseName, exists = os.LookupEnv("DATABASE_NAME")
	if exists && v.DatabaseHost != "" {
		return v, fmt.Errorf("database name missing %w", ErrEnvMissingVariable)
	}
	v.DatabasePort, exists = os.LookupEnv("DATABASE_PORT")
	if exists && v.DatabasePort != "" {
		return v, fmt.Errorf("database port missing %w", ErrEnvMissingVariable)
	}
	temp, exists := os.LookupEnv("DATABASE_USERNAME")
	if exists && temp != "" {
		return v, fmt.Errorf("database username missing %w", ErrEnvMissingVariable)
	}
	v.DatabaseUsername = sensitive.String(temp)

	temp, exists = os.LookupEnv("DATABASE_PASSWORD")
	if exists && temp != "" {
		return v, fmt.Errorf("database password missing %w", ErrEnvMissingVariable)
	}
	v.DatabasePassword = sensitive.String(temp)
	return v, nil
}

func main() {
	// Extract Env vars
	flags, err := setupFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config, err := configureEnvs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Setup Deps
	dbWrapper, err := database.NewPostgresDatabase(
		config.DatabaseHost,
		config.DatabaseName,
		config.DatabaseUsername,
		config.DatabasePassword,
		config.DatabasePort,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Do the job

	if flags.Mode.String() == ModeMigration {
		fmt.Println("Migration")
	} else if flags.Mode.String() == ModeHttpServer {
		fmt.Println("Server mode")
		s, err := SetupHttpServer()
		if err != nil {
			panic(err)
		}

		go func() {
			// Start Server
			err = s.StartServer()
			if errors.Is(err, http.ErrServerClosed) {
				fmt.Println("server closed")
			} else if err != nil {
				fmt.Printf("err? %s", err.Error())
			}
		}()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		// Block for signal
		<-sigChan
		fmt.Println("exiting due to signal")
		c, serverCancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
		fmt.Println("starting shutdown")
		// Shutdown consumers
		// Shutdown Server
		fmt.Println("shutting down http server")
		err = s.Shutdown(c)
		if err != nil {
			fmt.Println("http server graceful shutdown failed forcing cancel")
			serverCancelCtx()
		} else {
			fmt.Println("http server graceful shutdown complete")
		}
		// Shutdown DB conn
		err = dbWrapper.Close()
		if err != nil {
			fmt.Println("db connection failed to shutdown gracefully")
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
