package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Logger struct {
		Stdout        bool     `envconfig:"LOGGER_STDOUT"`
		FileLocation  string   `envconfig:"LOGGER_FILE_LOCATION"`
		FileMaxAge    int      `envconfig:"LOGGER_FILE_MAX_AGE"`
		Level         int8     `envconfig:"LOGGER_LEVEL"`
		Masking       bool     `envconfig:"LOGGER_MASKING"`
		MaskingParams []string `envconfig:"LOGGER_MASKING_PARAMS"`
	}
}

// LoadTest loads test config
func LoadTest() *Config {
	return load()
}

// LoadDefault loads default config from environment variables
func LoadDefault() *Config {
	return load()
}

// load config from environment variables
func load() *Config {
	var c Config

	_ = godotenv.Load() // Load .env file if it exists

	if err := envconfig.Process("", &c); err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	return &c
}
