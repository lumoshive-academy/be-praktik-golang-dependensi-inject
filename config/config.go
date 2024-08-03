package config

import "fmt"

type Config struct {
	Message string
}

type ConfigA Config
type ConfigB Config

func NewConfig() *ConfigA {
	config := &Config{
		Message: "Config 1",
	}

	return (*ConfigA)(config)
}

func NewConfigAlternative() *ConfigB {
	config := &Config{
		Message: "Config 1",
	}
	return (*ConfigB)(config)
}

// AppConfig adalah struct yang menyimpan konfigurasi aplikasi
type AppConfig struct {
	AppName string
	Port    int
	DBName  string
	Version int
}

// PrintConfig mencetak konfigurasi aplikasi
func (a *AppConfig) PrintConfig() {
	fmt.Printf("App Name: %s, Version: %v\n", a.AppName, a.Version)
}
