package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Dialect  string `mapstructure:"dialect"`
	URL      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Pool     int    `mapstructure:"pool"`
	IdlePool int    `mapstructure:"idle_pool"`
}

func setupDefaultsDatabase() {
	viper.SetDefault("db.dialect", "postgres")
	viper.SetDefault("db.url", "")
	viper.SetDefault("db.database", "tasks-management-with-go_development")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.user", "root")
	viper.SetDefault("db.password", "etu28")
	viper.SetDefault("db.pool", 0)
	viper.SetDefault("db.idle_pool", 2)
}
