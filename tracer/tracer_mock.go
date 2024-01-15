package tracer

import (
	"github.com/joaosoft/clean-infrastructure/domain"
	tracerConfig "github.com/joaosoft/clean-infrastructure/tracer/config"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/mock"
)

func NewTracerMock() *TracerMock {
	return &TracerMock{}
}

type TracerMock struct {
	mock.Mock
}

func (t *TracerMock) Name() string {
	args := t.Called()
	return args.Get(0).(string)
}

func (t *TracerMock) Start() error {
	args := t.Called()
	return args.Error(0)
}

func (t *TracerMock) Stop() error {
	args := t.Called()
	return args.Error(0)
}

func (t *TracerMock) Config() *tracerConfig.Config {
	args := t.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*tracerConfig.Config)
}

func (t *TracerMock) ConfigFile() string {
	args := t.Called()
	return args.Get(0).(string)
}

func (t *TracerMock) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	args := t.Called(operationName, opts)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(opentracing.Span)
}

func (t *TracerMock) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	args := t.Called(sm, format, carrier)
	return args.Error(0)
}

func (t *TracerMock) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	args := t.Called(format, carrier)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(opentracing.SpanContext), args.Error(1)
}

// WithAdditionalConfigType sets an additional config type
func (t *TracerMock) WithAdditionalConfigType(obj interface{}) domain.IApp {
	args := t.Called(obj)
	return args.Get(0).(domain.IApp)
}

// Started true if started
func (t *TracerMock) Started() bool {
	args := t.Called()
	return args.Get(0).(bool)
}
