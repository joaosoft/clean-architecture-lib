package table

import (
	"github.com/gocraft/dbr/v2"
	"github.com/stretchr/testify/mock"
)

func NewTableMock() *TableMock {
	return &TableMock{}
}

type TableMock struct {
	mock.Mock
}

func (t *TableMock) As(alias string) dbr.Builder {
	args := t.Called(alias)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(dbr.Builder)
}

func (t *TableMock) Build(d dbr.Dialect, buf dbr.Buffer) error {
	args := t.Called(d, buf)
	return args.Error(0)
}

func (t *TableMock) String() string {
	args := t.Called()
	return args.Get(0).(string)
}
