//go:build wireinject
// +build wireinject

package main

import (
	"golang-dependensi-inject/config"
	"golang-dependensi-inject/greeter"
	"golang-dependensi-inject/service"

	"github.com/google/wire"
)

func InitializMyService(name string) (*service.Service, error) {
	wire.Build(greeter.NewGreeter, service.NewService)
	return nil, nil
}

func InitializeServiceConfig() (*service.ServiceConfig, error) {
	wire.Build(config.NewConfig, config.NewConfigAlternative, service.NewServiceConfig)
	return nil, nil
}

// var SuperSet = wire.NewSet(
// 	greeter.NewGreeter,
// 	service.NewService,
// )
