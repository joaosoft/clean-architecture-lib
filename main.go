package main

import (
	controllers "clean-architecture/controllers/http/person"
	_ "clean-architecture/infrastructure"
	"clean-architecture/infrastructure/config/viper"
	"clean-architecture/infrastructure/http/server"
	models "clean-architecture/models/person"
	repositories "clean-architecture/repositories/person"
	"log"
)

func main() {
	repository, err := repositories.
		NewPersonRepository()
	if err != nil {
		panic(err)
	}

	personController := controllers.NewPersonController(
		models.NewPersonModel(repository))

	if err = server.
		New().
		WithConfigLoader(viper.NewViper()).
		WithLogger(log.Default()).
		WithController("person", personController).
		Start(); err != nil {
		panic(err)
	}
}
