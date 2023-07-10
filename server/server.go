// Package server provides functionality to easily set up an HTTTP server.
//
// The server holds all the clients it needs and they should be set up in the Create method.
//
// The HTTP routes and middleware are set up in the setupRouter method.
package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"ware-data-api/client/driver"
	"ware-data-api/config"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server holds the HTTP server, router, config and all clients.
type Server struct {
	Config    *config.Config
	HTTP      *http.Server
	DriverV45 *driver.Client
	Router    *mux.Router
}

// Create sets up the HTTP server, router and all clients.
// Returns an error if an error occurs.
func (s *Server) Create(ctx context.Context, config *config.Config) error {
	var drvClient driver.Client
	if err := drvClient.Init(config); err != nil {
		return fmt.Errorf("driver client: %w", err)
	}

	s.DriverV45 = &drvClient
	s.Config = config
	s.Router = mux.NewRouter()
	s.HTTP = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Port),
		Handler: s.Router,
	}

	s.setupRoutes()

	return nil
}

// Serve tells the server to start listening and serve HTTP requests.
// It also makes sure that the server gracefully shuts down on exit.
// Returns an error if an error occurs.
func (s *Server) Serve(ctx context.Context) error {
	idleConnsClosed := make(chan struct{}) // this is used to signal that we can not exit
	go func(ctx context.Context, s *http.Server) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

		<-stop

		log.Info("Shutdown signal received")

		if err := s.Shutdown(ctx); err != nil {
			log.Error(err.Error())
		}

		close(idleConnsClosed) // call close to say we can now exit the function
	}(ctx, s.HTTP)

	log.Infof("Ready at: %s", s.Config.Port)

	if err := s.HTTP.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("unexpected server error: %w", err)
	}
	<-idleConnsClosed // this will block until close is called

	return nil
}
