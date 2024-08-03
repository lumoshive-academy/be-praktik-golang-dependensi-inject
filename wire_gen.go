// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"golang-dependensi-inject/config"
	"golang-dependensi-inject/greeter"
	"golang-dependensi-inject/notification"
	"golang-dependensi-inject/service"
	"golang-dependensi-inject/storage"
	"io"
	"os"
)

// Injectors from wire.go:

func InitializMyService(name string) (*service.Service, error) {
	greeterGreeter := greeter.NewGreeter(name)
	serviceService, err := service.NewService(greeterGreeter)
	if err != nil {
		return nil, err
	}
	return serviceService, nil
}

func InitializeServiceConfig() *service.ServiceConfig {
	configA := config.NewConfig()
	configB := config.NewConfigAlternative()
	serviceServiceConfig := service.NewServiceConfig(configA, configB)
	return serviceServiceConfig
}

// InitializeCachingData menginisialisasi storage menggunakan InMemoryStorageSet
func InitializeCachingData() (storage.Storage, error) {
	cachingData := service.NewCachingData()
	return cachingData, nil
}

// InitializeDatabaseStorage menginisialisasi storage menggunakan DatabaseStorageSet
func InitializeDatabaseStorage() (storage.Storage, error) {
	databaseStorage := service.NewDatabaseStorage()
	return databaseStorage, nil
}

// InitializeNotifier menginisialisasi Notifier dengan ketergantungan yang diperlukan
func InitializeNotifier() (*service.Notifier, error) {
	emailService := notification.NewEmailService()
	smsService := notification.NewSMSService()
	notifConfig := config.NewNotifConfig()
	notifier := &service.Notifier{
		EmailService: emailService,
		SMSService:   smsService,
		NotifConfig:  notifConfig,
	}
	return notifier, nil
}

// InitializeAppConfig menginisialisasi AppConfig dengan nilai konstan
func InitializeAppConfig() (*config.AppConfig, error) {
	string2 := _wireStringValue
	int2 := _wireIntValue
	appConfig := &config.AppConfig{
		AppName: string2,
		Version: int2,
	}
	return appConfig, nil
}

var (
	_wireStringValue = "MyApp"
	_wireIntValue    = 1
)

func injectReader() io.Reader {
	reader := _wireFileValue
	return reader
}

var (
	_wireFileValue = os.Stdin
)

func InitializeDatabase() *storage.Database {
	appConfig := _wireAppConfigValue
	string2 := appConfig.DBName
	database := storage.NewDatabase(string2)
	return database
}

var (
	_wireAppConfigValue = AppConfig
)

func InitializeDatabaseWithCleanUp() (*storage.Database, func(), error) {
	appConfig := _wireAppConfigValue
	string2 := appConfig.DBName
	database, cleanup, err := storage.NewDatabaseWithCleanUp(string2)
	if err != nil {
		return nil, nil, err
	}
	return database, func() {
		cleanup()
	}, nil
}

// wire.go:

var myservice = wire.NewSet(greeter.NewGreeter, service.NewService)

var serviceConfig = wire.NewSet(config.NewConfig, config.NewConfigAlternative, service.NewServiceConfig)

// cachingDataSet menghubungkan Storage dengan InMemoryStorage
var cachingDataSet = wire.NewSet(service.NewCachingData, wire.Bind(new(storage.Storage), new(*storage.CachingData)))

// DatabaseStorageSet menghubungkan Storage dengan DatabaseStorage
var databaseStorageSet = wire.NewSet(service.NewDatabaseStorage, wire.Bind(new(storage.Storage), new(*storage.DatabaseStorage)))

var notifierSet = wire.NewSet(notification.NewEmailService, notification.NewSMSService, config.NewNotifConfig)

var AppConfig = config.AppConfig{
	AppName: "MyApp",
	Port:    8080,
	DBName:  "mydatabase",
}

var ConfigSet = wire.NewSet(wire.Value(AppConfig), wire.FieldsOf(new(config.AppConfig), "DBName"))

var DatabaseSet = wire.NewSet(
	ConfigSet, storage.NewDatabase,
)

var DatabaseSetWithCleanUp = wire.NewSet(
	ConfigSet, storage.NewDatabaseWithCleanUp,
)
