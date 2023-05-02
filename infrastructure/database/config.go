package database

type Database struct {
	Driver     string `mapstructure:"driver"`
	DataSource string `mapstructure:"data_source"`
}
