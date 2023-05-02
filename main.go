package main

import (
	controllers "clean-architecture/controllers/http/person"
	appHttp "clean-architecture/infrastructure/app/http"
	"clean-architecture/infrastructure/config/viper"
	"clean-architecture/infrastructure/database/postgres"
	models "clean-architecture/models/person"
	repositories "clean-architecture/repositories/person"
	"log"
)

func main() {
	app := appHttp.New()

	config, err := viper.NewViper().Load()
	if err != nil {
		panic(err)
	}

	db, err := postgres.NewConnection(config.Database.Driver, config.Database.DataSource)
	if err != nil {
		panic(err)
	}

	personRepo, err := repositories.NewPersonRepository(app)
	if err != nil {
		panic(err)
	}

	personController := controllers.NewPersonController(app, models.NewPersonModel(app, personRepo))

	app.WithConfig(config).
		WithLogger(log.Default()).
		WithDb(db).
		WithController(personController)

	if err = app.Start(); err != nil {
		panic(err)
	}
}
