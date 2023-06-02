package config

import (
	"github.com/joaosoft/clean-architecture/infrastructure/app"
	"github.com/joaosoft/clean-architecture/infrastructure/database"
)

type Config struct {
	Http     app.Http          `mapstructure:"http"`
	Database database.Database `mapstructure:"database"`
}
