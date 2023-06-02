package main

import (
	controllers "github.com/joaosoft/clean-architecture/controllers/http/person"
	appHttp "github.com/joaosoft/clean-architecture/infrastructure/app/http"
	"github.com/joaosoft/clean-architecture/infrastructure/config/viper"
	"github.com/joaosoft/clean-architecture/infrastructure/database/postgres"
	models "github.com/joaosoft/clean-architecture/models/person"
	repositories "github.com/joaosoft/clean-architecture/repositories/person"
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
