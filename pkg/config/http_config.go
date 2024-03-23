package config

type HttpConfig struct {
	Port string `json:"port" mapstructure:"port" validate:"required"`
}
