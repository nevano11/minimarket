package config

import "github.com/nevano11/minimarket/pkg/config"

type Config struct {
	Database config.DatabaseConfig `json:"database" mapstructure:"database" validate:"required"`
	Server   config.HttpConfig     `json:"server"   mapstructure:"server"   validate:"required"`
}

func Load(configPath string) (*Config, error) {
	cfg := &Config{}

	err := config.LoadFromFile(configPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
