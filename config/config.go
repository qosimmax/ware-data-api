// Package config handles environment variables.
package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Config contains environment variables.
type Config struct {
	Port    string `envconfig:"PORT" default:"8000"`
	UDPPort string `envconfig:"UDP_PORT" default:"1111"`
	FromIP  string `envconfig:"FROM_IP" default:"192.168.0.1"`
	ToIP    string `envconfig:"TO_IP" default:"192.168.0.255"`
}

// LoadConfig reads environment variables and populates Config.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	var c Config

	err := envconfig.Process("", &c)

	return &c, err
}
