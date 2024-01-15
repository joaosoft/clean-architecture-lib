package config

import (
	"github.com/joaosoft/clean-infrastructure/tracer/jaeger/config"
)

// Config configurations for the tracer
type Config struct {
	// Enabled
	Enabled bool `yaml:"enabled"`
	// Strategy
	Strategy string `yaml:"strategy"`
	// Jaeger
	Jaeger config.Jaeger `json:"jaeger"`
	// Additional Config
	AdditionalConfig interface{} `yaml:"additionalConfig"`
}
