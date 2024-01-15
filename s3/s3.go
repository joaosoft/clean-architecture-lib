package s3

import (
	"context"

	"github.com/joaosoft/clean-infrastructure/domain/message"

	"github.com/joaosoft/clean-infrastructure/errors"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joaosoft/clean-infrastructure/config"
	"github.com/joaosoft/clean-infrastructure/domain"
	s3Config "github.com/joaosoft/clean-infrastructure/s3/config"
)

// S3 service
type S3 struct {
	// Name
	name string
	// App
	app domain.IApp
	// Configuration
	config *s3Config.Config
	// Client
	s3Client *s3.Client
	// Additional Config Type
	additionalConfigType interface{}
	// Started
	started bool
}

const (
	// configFile sqs configuration file
	configFile = "s3.yaml"
)

// New creates a new sqs service
func New(app domain.IApp, config *s3Config.Config) *S3 {
	s := &S3{
		app:  app,
		name: "s3",
	}

	if config != nil {
		s.config = config
	}

	return s
}

// Name gets the service name
func (s *S3) Name() string {
	return s.name
}

// Start starts the sqs service
func (s *S3) Start() (err error) {
	// Load config File
	if s.config == nil {
		s.config = &s3Config.Config{}
		s.config.AdditionalConfig = s.additionalConfigType
		if err = config.Load(s.ConfigFile(), s.config); err != nil {
			err = errors.ErrorLoadingConfigFile().Formats(s.ConfigFile(), err)
			message.ErrorMessage(s.Name(), err)
			return err
		}
	}
	// Init s3 config default
	cfg, _ := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithSharedConfigProfile(
			s.Config().Role,
		),
	)
	// Set Region
	cfg.Region = s.config.Region
	// Init S3 client
	s.s3Client = s3.NewFromConfig(cfg)

	s.started = true

	return nil
}

// Client stops the s3 service
func (s *S3) Client() *s3.Client {
	return s.s3Client
}

// Stop stops the s3 service
func (s *S3) Stop() error {
	if !s.started {
		return nil
	}
	s.started = false
	return nil
}

// Config gets the service configuration
func (s *S3) Config() *s3Config.Config {
	return s.config
}

// ConfigFile gets the configuration file
func (s *S3) ConfigFile() string {
	return configFile
}

// WithAdditionalConfigType sets an additional config type
func (s *S3) WithAdditionalConfigType(obj interface{}) domain.IS3 {
	s.additionalConfigType = obj
	return s
}

// Started true if started
func (s *S3) Started() bool {
	return s.started
}
