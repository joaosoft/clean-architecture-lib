package main

import (
	controllers "clean-architecture/controllers/http/person"
	_ "clean-architecture/infrastructure"
	"clean-architecture/infrastructure/config/viper"
	"clean-architecture/infrastructure/http"
	models "clean-architecture/models/person"
	repositories "clean-architecture/repositories/person"
	"log"
)

func main() {
	cfg, err := viper.Load()
	if err != nil {
		panic(err)
	}

	repository, err := repositories.NewPersonRepository(cfg)
	if err != nil {
		panic(err)
	}

	controller := controllers.NewPersonController(models.NewPersonModel(repository))

	if err = http.
		New(cfg.Http.Port).
		WithController(controller).
		WithLogger(log.Default()).
		Start(); err != nil {
		panic(err)
	}
}
