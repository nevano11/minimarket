package config

type DatabaseConfig struct {
	Host     string `json:"host"     mapstructure:"host"     validate:"required"`
	Port     string `json:"port"     mapstructure:"port"     validate:"required"`
	Username string `json:"username" mapstructure:"username" validate:"required"`
	Password string `json:"password" mapstructure:"password" validate:"required"`
	DbName   string `json:"dbName"   mapstructure:"dbName"   validate:"required"`
	SslMode  string `json:"sslMode"  mapstructure:"sslMode"  validate:"required"`
}
