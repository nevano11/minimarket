package config

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	ErrConfigPathIsEmpty = fmt.Errorf("config path is empty")
)

func LoadFromFile(configPath string, config any) error {
	if configPath != "" {
		viper.SetConfigFile(configPath)
		log.Infof("Reading config from %s", configPath)
	} else {
		return ErrConfigPathIsEmpty
	}

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("couldn't read config: %w", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("couldn't unmarshal: %w", err)
	}

	loaded, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("couldn't marshal: %w", err)
	}

	log.Infof("Loaded config is: %s", string(loaded))

	return nil
}
