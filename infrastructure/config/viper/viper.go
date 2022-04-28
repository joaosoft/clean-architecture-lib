package viper

import (
	"clean-architecture/infrastructure/config"

	"github.com/spf13/viper"
)

func Load() (_ *config.Config, err error) {
	viper.AddConfigPath("./config")

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &config.Config{}
	if err = viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
