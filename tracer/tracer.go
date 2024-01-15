package tracer

import (
	"fmt"
	"io"

	"github.com/joaosoft/clean-infrastructure/domain/message"

	errorCodes "github.com/joaosoft/clean-infrastructure/errors"

	"github.com/opentracing/opentracing-go"

	"github.com/joaosoft/clean-infrastructure/config"
	"github.com/joaosoft/clean-infrastructure/domain"
	tracerConfig "github.com/joaosoft/clean-infrastructure/tracer/config"
)

// Tracer service
type Tracer struct {
	// Name
	name string
	// Configuration
	config *tracerConfig.Config
	// App
	app domain.IApp
	// strategy
	strategy strategy
	// Client
	opentracing.Tracer
	// CLoser
	closer io.Closer
	// Additional Config Type
	additionalConfigType interface{}
	// Started
	started bool
}

const (
	// configFile tracer configuration file
	configFile = "tracer.yaml"
)

// New creates a new tracer service
func New(app domain.IApp, config *tracerConfig.Config) *Tracer {
	tracer := &Tracer{
		name: "Tracer",
		app:  app,
	}

	if config != nil {
		tracer.config = config
	}

	return tracer
}

// Name gets the service name
func (t *Tracer) Name() string {
	s := t.strategy
	if s == "" {
		s = domain.NotApplicable
	}
	return fmt.Sprintf("%s [%s]", t.name, s)
}

// Start starts the tracer service
func (t *Tracer) Start() (err error) {
	if t.config == nil {
		t.config = &tracerConfig.Config{}
		t.config.AdditionalConfig = t.additionalConfigType
		if err = config.Load(t.ConfigFile(), t.config); err != nil {
			err = errorCodes.ErrorLoadingConfigFile().Formats(t.ConfigFile(), err)
			message.ErrorMessage(t.Name(), err)
			return err
		}
	}

	// tracer
	if t.config.Enabled {
		if t.Tracer == nil {
			t.strategy = NewStrategy(t.config.Strategy)
			if t.Tracer, t.closer, err = t.strategy.Handle(t.app.Config().Name, t.config); err != nil {
				return err
			}
		}

		if t.Tracer != nil {
			// set the global tracer
			opentracing.SetGlobalTracer(t.Tracer)
		}
	}

	t.started = true

	return nil
}

// Stop stops the tracer service
func (t *Tracer) Stop() error {
	if !t.started {
		return nil
	}

	if t.closer != nil {
		if err := t.closer.Close(); err != nil {
			return err
		}
	}

	t.Tracer = opentracing.NoopTracer{}
	opentracing.SetGlobalTracer(t.Tracer)

	t.started = false
	return nil
}

// Config gets the configurations
func (t *Tracer) Config() *tracerConfig.Config {
	return t.config
}

// ConfigFile gets the configuration file
func (t *Tracer) ConfigFile() string {
	return configFile
}

// WithAdditionalConfigType sets an additional config type
func (t *Tracer) WithAdditionalConfigType(obj interface{}) domain.ITracer {
	t.additionalConfigType = obj
	return t
}

// Started true if started
func (t *Tracer) Started() bool {
	return t.started
}
