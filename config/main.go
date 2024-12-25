package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port         string
		Debug        bool
		AllowOrigins []string
	}
}

func NewConfig() *Config {

	return &Config{}
}

func (cfg *Config) LoadConfig() error {

	viper.SetConfigFile("temp/config.yml")

	cfg.setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unpack config: %v", err)
	}

	return nil
}
