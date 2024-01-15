package jaeger

import (
	"fmt"
	"io"

	jaegerConfig "github.com/joaosoft/clean-infrastructure/tracer/jaeger/config"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"

	"github.com/opentracing/opentracing-go"
)

// NewJaegerTracer created a new jaeger tracer
func NewJaegerTracer(appName string, cfg jaegerConfig.Jaeger) (opentracing.Tracer, io.Closer, error) {
	var collectorHostPort string
	if cfg.CollectorHostPort != "" {
		collectorHostPort = fmt.Sprintf("http://%s/api/traces", cfg.CollectorHostPort)
	}

	configuration := &config.Configuration{
		ServiceName: appName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           cfg.Log,
			LocalAgentHostPort: cfg.AgentHostPort,
			CollectorEndpoint:  collectorHostPort,
		},
	}

	return configuration.NewTracer(config.Logger(jaeger.StdLogger))
}
