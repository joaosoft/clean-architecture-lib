package http

import (
	routes "clean-architecture/api/http"
	domain "clean-architecture/domain/person"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	WithController(controller domain.IPersonController) *Server
	WithLogger(logger domain.ILogger) *Server
	Start() error
	Stop() error
}

type Server struct {
	App        *http.Server
	Router     *gin.Engine
	Port       int
	controller domain.IPersonController
	logger     domain.ILogger
}

func New(port int) *Server {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	server := &Server{
		App: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		Router: router,
		Port:   port,
	}

	router.Use(gin.Recovery())

	return server
}

func (s *Server) WithLogger(logger domain.ILogger) *Server {
	s.logger = logger
	return s
}

func (s *Server) WithController(controller domain.IPersonController) *Server {
	s.controller = controller
	routes.Register(s.Router, controller)
	return s
}

func (s *Server) Start() (err error) {
	return s.App.ListenAndServe()
}

func (s *Server) Stop() (err error) {
	return s.App.Shutdown(context.Background())
}
