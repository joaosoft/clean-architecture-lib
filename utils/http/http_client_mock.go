package http

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type HttpClientMock struct {
	mock.Mock
}

func NewHttpClientMock() *HttpClientMock {
	return &HttpClientMock{}
}

func (c *HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	args := c.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*http.Response), args.Error(1)
}
