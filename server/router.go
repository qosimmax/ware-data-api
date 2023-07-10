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
	//api.HandleFunc("/send-sms", handler.CreateSms(s.DB, s.Storage, s.PubSub)).Methods(http.MethodPost).Name("SendSms")
	//api.HandleFunc("/sms-report", handler.GenerateSmsReport(s.DB, s.Minio)).Methods(http.MethodPost).Name("SmsReport")
	_ = api
}