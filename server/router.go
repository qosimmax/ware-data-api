package server

import (
	"net/http"
	"ware-data-api/server/internal/handler"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const v1API string = "/ware-data-api/api/v1"

// setupRoutes - the root route function.
func (s *Server) setupRoutes() {
	s.Router.Handle("/metrics", promhttp.Handler()).Name("Metrics")
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods(http.MethodGet).Name("Health")

	api := s.Router.PathPrefix(v1API).Subrouter()
	api.HandleFunc("/devices", handler.GetDevices(s.DriverV45)).Methods(http.MethodGet).Name("GetDevices")
	api.HandleFunc("/wares", handler.WaresUpload(s.DriverV45)).Methods(http.MethodPost).Name("WaresUpload")

	go handler.LoadDevices(s.DriverV45)

}
