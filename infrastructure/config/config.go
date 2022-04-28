package config

import (
	"clean-architecture/infrastructure/database"
	"clean-architecture/infrastructure/http"
)

type Config struct {
	Http     http.Http         `mapstructure:"http"`
	Database database.Database `mapstructure:"database"`
}
