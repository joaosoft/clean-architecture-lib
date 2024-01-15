package logging

import (
	"github.com/joaosoft/clean-infrastructure/domain"
	"github.com/joaosoft/clean-infrastructure/logger/config"
	"github.com/joaosoft/clean-infrastructure/utils/errors"
	"github.com/stretchr/testify/mock"
)

func NewLoggingMock() *LoggingMock {
	return &LoggingMock{}
}

type LoggingMock struct {
	mock.Mock
}

func (l *LoggingMock) Do(err error, info ...*domain.LoggerInfo) {
	var optsList = []interface{}{err}
	for _, v := range info {
		optsList = append(optsList, v)
	}

	_ = l.Called(optsList...)
}

func (l *LoggingMock) Multi(err []error, info ...*domain.LoggerInfo) {
	var optsList = []interface{}{err}
	for _, v := range info {
		optsList = append(optsList, v)
	}

	_ = l.Called(optsList...)
}

func (l *LoggingMock) Frontend(error string, level errors.Level, fe *domain.Frontend) {
	l.Called(error, level, fe)
}

func (l *LoggingMock) Init(cgf config.Config) domain.ILogging {
	args := l.Called(cgf)
	return args.Get(0).(domain.ILogging)
}
