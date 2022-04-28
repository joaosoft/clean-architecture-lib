package main

import (
	"clean-architecture/controllers"
	"clean-architecture/infrastructure/database"
	"clean-architecture/infrastructure/http"
	"clean-architecture/models"
	"clean-architecture/repositories"

	"github.com/spf13/viper"
)

func main() {
	var err error
	if err = loadConfigs(); err != nil {
		panic(err)
	}

	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	if err = http.New(viper.GetInt("http_port")).
		WithController(controllers.NewController(models.NewModel(repositories.NewRepository(db)))).
		Start(); err != nil {
		panic(err)
	}
}

func loadConfigs() error {
	viper.AddConfigPath("./config")
	return viper.ReadInConfig()
}
