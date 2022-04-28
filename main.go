package main

import (
	"clean-architecture/controllers"
	"clean-architecture/infrastructure/database"
	"clean-architecture/infrastructure/http"
	"clean-architecture/models"
	"clean-architecture/repositories"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	if err = http.New(8081).
		WithController(controllers.NewController(models.NewModel(repositories.NewRepository(db)))).
		Start(); err != nil {
		panic(err)
	}
}
