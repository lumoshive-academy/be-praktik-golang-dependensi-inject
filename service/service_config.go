package service

import "golang-dependensi-inject/config"

type ServiceConfig struct {
	Config *config.Config
}

func NewServiceConfig(cfg *config.Config) *ServiceConfig {
	return &ServiceConfig{Config: cfg}
}
