package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort       string
	AuthServiceAddr  string
	TasksServiceAddr string
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv() // Allows to override config values with environment variables

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
