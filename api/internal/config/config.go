package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `envconfig:"PORT" default:"3000"`
}

func Load() (*Config, error) {
	var c Config
	return &c, envconfig.Process("cf", &c)
}
