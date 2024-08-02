package config

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
