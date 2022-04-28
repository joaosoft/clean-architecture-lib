package config

import (
	"clean-architecture/infrastructure/database"
	"clean-architecture/infrastructure/http"

	"github.com/spf13/viper"
)

type Config struct {
	Http     http.Http         `mapstructure:"http"`
	Database database.Database `mapstructure:"database"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Load() (err error) {
	viper.AddConfigPath("./config")

	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	if err = viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}
