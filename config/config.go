package config

type Config struct {
	Message string
}

func NewConfig() *Config {
	return &Config{Message: "Config 1"}
}

func NewConfigAlternative() *Config {
	return &Config{Message: "Config 2"}
}
