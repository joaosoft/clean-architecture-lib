package datatable

import (
	databaseConfig "github.com/joaosoft/clean-infrastructure/database/config"
	datatableConfig "github.com/joaosoft/clean-infrastructure/datatable/config"
	dtDatabase "github.com/joaosoft/clean-infrastructure/datatable/database"
	dtElastic "github.com/joaosoft/clean-infrastructure/datatable/elastic_search"
	"github.com/joaosoft/clean-infrastructure/domain"
)

type Datatable struct {
	// Name
	name string
	// Configuration
	config *databaseConfig.Config //nolint:all
	// App
	app domain.IApp
	// Max results per pages
	maxPageResultLimit int
	// Log functions
	logFunction logFunction //nolint:all
	// Services
	services services
	// Additional Config Type
	additionalConfigType interface{}
	// Started
	started bool
}

type services struct {
	database dtDatabase.IDatabase
	elastic  dtElastic.IElastic
}

type logFunction func(err error) error //nolint:all

// New creates a new database
func New(app domain.IApp, config *datatableConfig.Config) *Datatable {
	datatable := &Datatable{
		name:               "Datatable",
		maxPageResultLimit: dtDatabase.MaxPageResultLimit,
		app:                app,
	}

	if config != nil {
		datatable.maxPageResultLimit = config.MaxPageResultLimit
	}

	return datatable
}

// Name name of the service
func (d *Datatable) Name() string {
	return d.name
}

// Start starts the service
func (d *Datatable) Start() error {

	if d.app.ElasticSearch() != nil { //nolint:all
		//Not implemented yet due not any valid use case to use it
	}

	if d.app.Database() != nil {
		d.services.database = dtDatabase.New(
			dtDatabase.Client{
				Reader: d.app.Database().Read(),
				Writer: d.app.Database().Write(),
			},
			d.app.Logger().DBLog,
			d.maxPageResultLimit,
		)
	}

	d.started = true

	return nil
}

// Stop stops the service
func (d *Datatable) Stop() error {
	if !d.started {
		return nil
	}
	d.started = false
	return nil
}

// Database datatable
func (d *Datatable) Database() dtDatabase.IDatabase {
	return d.services.database
}

// Database datatable
func (d *Datatable) Elastic() dtElastic.IElastic {
	return d.services.elastic
}

// WithAdditionalConfigType sets an additional config type
func (d *Datatable) WithAdditionalConfigType(obj interface{}) domain.IDatatable {
	d.additionalConfigType = obj
	return d
}

// Started true if started
func (d *Datatable) Started() bool {
	return d.started
}
