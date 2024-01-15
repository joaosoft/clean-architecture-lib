package tracer

import (
	"io"

	"github.com/joaosoft/clean-infrastructure/tracer/jaeger"

	"github.com/joaosoft/clean-infrastructure/tracer/config"
	"github.com/opentracing/opentracing-go"
)

// strategy type
type strategy string

// Strategies
const (
	StrategyJaeger strategy = "jaeger"
)

// NewStrategy creates a new strategy
func NewStrategy(s string) strategy {
	return strategy(s)
}

// Handle gets the strategy implementation
func (s *strategy) Handle(appName string, cfg *config.Config) (opentracing.Tracer, io.Closer, error) {
	switch *s {
	case StrategyJaeger:
		return jaeger.NewJaegerTracer(appName, cfg.Jaeger)
	default:
		return nil, nil, nil
	}
}

// String prints the strategy
func (s *strategy) String() string {
	return string(*s)
}
