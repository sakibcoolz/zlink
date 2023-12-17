package config

import (
	"os"
	"strconv"

	"go.uber.org/zap"
)

type ServiceConfig struct {
	Host string
	Port int
}

type Config struct {
	Service ServiceConfig
}

func NewConfig(log *zap.Logger) *Config {
	var config Config
	log.Info("Fetching service details")

	if config.Service.Host = os.Getenv("SERVICEHOST"); config.Service.Host == "" {
		log.Fatal("SERVICEHOST environment variable not set")
	}

	if config.Service.Port, _ = strconv.Atoi(os.Getenv("SERVICEPORT")); config.Service.Port == 0 {
		log.Fatal("SERVICEPORT environment variable not set")
	}

	return &config
}

func (config *Config) GetServiceConfig() ServiceConfig {
	return config.Service
}
