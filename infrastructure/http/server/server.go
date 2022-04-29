package server

import (
	routes "clean-architecture/api/http"
	domain "clean-architecture/domain"
	"clean-architecture/infrastructure/config"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	App          *http.Server
	Router       *gin.Engine
	configLoader domain.IConfig
	config       *config.Config
	logger       domain.ILogger
	controllers  map[string]domain.IController
}

func New() *Server {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	server := &Server{
		App: &http.Server{
			Handler: router,
		},
		Router: router,
	}

	router.Use(gin.Recovery())

	return server
}

func (s *Server) WithLogger(logger domain.ILogger) *Server {
	s.logger = logger
	return s
}

func (s *Server) WithConfigLoader(configLoader domain.IConfig) *Server {
	s.configLoader = configLoader
	return s
}

func (s *Server) WithController(name string, controller domain.IController) *Server {
	s.controllers[name] = controller
	return s
}

func (s *Server) Setup() (err error) {
	if s.config, err = s.configLoader.Load(); err != nil {
		return err
	}

	s.App.Addr = fmt.Sprintf(":%d", s.config.Http.Port)

	for _, c := range s.controllers {
		if c != nil {
			return c.Setup(s.config, s.logger)
		}
	}

	return routes.Register(s.Router, s.controllers)
}

func (s *Server) Start() (err error) {
	if err = s.Setup(); err != nil {
		return err
	}
	return s.App.ListenAndServe()
}

func (s *Server) Stop() (err error) {
	return s.App.Shutdown(context.Background())
}
