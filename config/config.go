package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-search-scrubber-api
type Config struct {
	BindAddr string `envconfig:"BIND_ADDR"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	cfg = &Config{
		BindAddr: ":3333",
	}

	return cfg, envconfig.Process("", cfg)
}
