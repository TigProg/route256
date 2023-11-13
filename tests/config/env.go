//go:build integration
// +build integration

package config

import (
	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "QA"

type Config struct {
	Host       string `split_words:"true" default:"localhost:8081"`
	DbHost     string `split_words:"true" default:"localhost"`
	DbPort     int    `split_words:"true" default:"5432"`
	DbUser     string `split_words:"true" default:"user"`
	DbPassword string `split_words:"true" default:"password"`
	DbName     string `split_words:"true" default:"bus_booking"`
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process(envPrefix, cfg)
	return cfg, err
}
