package service

import "golang-dependensi-inject/config"

type ServiceConfig struct {
	ConfigA *config.ConfigA
	ConfigB *config.ConfigB
}

func NewServiceConfig(cfga *config.ConfigA, cfgb *config.ConfigB) *ServiceConfig {
	return &ServiceConfig{ConfigA: cfga, ConfigB: cfgb}
}
