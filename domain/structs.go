package domain

import "clean-architecture/infrastructure/config"

type App struct {
	Config *config.Config
	Logger ILogger
}
