package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Address string `mapstructure: "address"`
}

func setupDefaultServer() {
	viper.SetDefault("server.address", "localhost:8080")
}
