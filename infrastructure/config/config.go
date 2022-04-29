package config

import (
	"clean-architecture/infrastructure/app"
	"clean-architecture/infrastructure/database"
)

type Config struct {
	Http     app.Http          `mapstructure:"http"`
	Database database.Database `mapstructure:"database"`
}
