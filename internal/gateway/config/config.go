package config

import (
	"fmt"

	"github.com/MiralSNk/go-micro-blog/internal/config"
)

type Config struct {
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	LogLevel string `yaml:"log_level"`
}

// LoadConfig создает конфигурацию
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := config.Load("gateway/config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf("failed: config not found %v", err)
	}

	return cfg, nil
}
