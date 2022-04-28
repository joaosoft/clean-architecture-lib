package main

import (
	"clean-architecture/controllers"
	"clean-architecture/infrastructure/config"
	"clean-architecture/infrastructure/http"
	"clean-architecture/models"
	"clean-architecture/repositories"
)

func main() {
	cfg := config.New()
	if err := cfg.Load(); err != nil {
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
