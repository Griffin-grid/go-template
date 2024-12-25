package config

import "github.com/spf13/viper"

func (cfg *Config) setDefaults() {

	// Here we set config variables default values

	viper.SetDefault("APP.PORT", "8080")
	viper.SetDefault("APP.DEBUG", true)
}
