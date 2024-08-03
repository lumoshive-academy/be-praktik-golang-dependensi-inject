//go:build wireinject
// +build wireinject

package main

import (
	"golang-dependensi-inject/config"
	"golang-dependensi-inject/greeter"
	"golang-dependensi-inject/notification"
	"golang-dependensi-inject/service"
	"golang-dependensi-inject/storage"
	"io"
	"os"

	"github.com/google/wire"
)

var myservice = wire.NewSet(
	greeter.NewGreeter,
	service.NewService,
)

func InitializMyService(name string) (*service.Service, error) {
	wire.Build(myservice)
	return nil, nil
}

var serviceConfig = wire.NewSet(
	config.NewConfig,
	config.NewConfigAlternative,
	service.NewServiceConfig,
)

func InitializeServiceConfig() *service.ServiceConfig {
	wire.Build(serviceConfig)
	return nil
}

// cachingDataSet menghubungkan Storage dengan InMemoryStorage
var cachingDataSet = wire.NewSet(
	service.NewCachingData,
	wire.Bind(new(storage.Storage), new(*storage.CachingData)),
)

// DatabaseStorageSet menghubungkan Storage dengan DatabaseStorage
var databaseStorageSet = wire.NewSet(
	service.NewDatabaseStorage,
	wire.Bind(new(storage.Storage), new(*storage.DatabaseStorage)),
)

// InitializeCachingData menginisialisasi storage menggunakan InMemoryStorageSet
func InitializeCachingData() (storage.Storage, error) {
	wire.Build(cachingDataSet)
	return nil, nil
}

// InitializeDatabaseStorage menginisialisasi storage menggunakan DatabaseStorageSet
func InitializeDatabaseStorage() (storage.Storage, error) {
	wire.Build(databaseStorageSet)
	return nil, nil
}

var notifierSet = wire.NewSet(
	notification.NewEmailService,
	notification.NewSMSService,
	config.NewNotifConfig,
)

// InitializeNotifier menginisialisasi Notifier dengan ketergantungan yang diperlukan
func InitializeNotifier() (*service.Notifier, error) {
	wire.Build(
		notifierSet,
		wire.Struct(new(service.Notifier), "*"),
	)
	return nil, nil
}

func injectReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// InitializeAppConfig menginisialisasi AppConfig dengan nilai konstan
func InitializeAppConfig() (*config.AppConfig, error) {
	wire.Build(
		wire.Value("MyApp"),
		wire.Value(1),
		wire.Struct(new(config.AppConfig), "AppName", "Version"),
	)
	return nil, nil
}
