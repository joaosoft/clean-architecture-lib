package buffer

import (
	"github.com/stretchr/testify/mock"
)

func NewBufferMock() *BufferMock {
	return &BufferMock{}
}

type BufferMock struct {
	mock.Mock
}

func (b *BufferMock) WriteValue(v ...interface{}) error {
	args := b.Called(v...)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (b *BufferMock) Value() []interface{} {
	args := b.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]interface{})
}
