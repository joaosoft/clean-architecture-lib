package main

import (
	controllers "clean-architecture/controllers/http"
	_ "clean-architecture/infrastructure"
	"clean-architecture/infrastructure/config/viper"
	"clean-architecture/infrastructure/http"
	"clean-architecture/models"
	"clean-architecture/repositories"
)

func main() {
	cfg, err := viper.Load()
	if err != nil {
		panic(err)
	}

	repository, err := repositories.NewRepository(cfg)
	if err != nil {
		panic(err)
	}

	controller := controllers.NewController(models.NewModel(repository))

	if err = http.
		New(cfg.Http.Port).
		WithController(controller).
		Start(); err != nil {
		panic(err)
	}
}
