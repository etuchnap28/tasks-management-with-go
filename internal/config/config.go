package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Database *DatabaseConfig `mapstructure:"db"`
	Server   *ServerConfig   `mapstructure:"server"`
}

func ReadAppConfig() *Config {
	viper.SetConfigName("app_config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("app_config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	setupDefaults()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return &config
}
