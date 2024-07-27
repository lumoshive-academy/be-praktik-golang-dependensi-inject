//go:build wireinject
// +build wireinject

package main

import (
	"golang-dependensi-inject/greeter"
	"golang-dependensi-inject/service"

	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	greeter.NewGreeter,
	service.NewService,
)

func InitializMyService() (*service.Service, error) {
	wire.Build(SuperSet)
	return nil, nil
}
