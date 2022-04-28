package main

import (
	"clean-architecture/config"
	"clean-architecture/controllers"
	"clean-architecture/infrastructure/database"
	"clean-architecture/infrastructure/http"
	"clean-architecture/models"
	"clean-architecture/repositories"

	"github.com/spf13/viper"
)

func main() {
	var err error
	if err = config.Load(); err != nil {
		panic(err)
	}

	db, err := database.NewDatabase(
		viper.GetString(config.DatabaseDriverConfigKey),
		viper.GetString(config.DatabaseDataSourceConfigKey),
	)
	if err != nil {
		panic(err)
	}

	if err = http.New(viper.GetInt(config.HttpPortConfigKey)).
		WithController(controllers.NewController(models.NewModel(repositories.NewRepository(db)))).
		Start(); err != nil {
		panic(err)
	}
}
