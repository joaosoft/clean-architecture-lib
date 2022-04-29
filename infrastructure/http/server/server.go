package server

import (
	routes "clean-architecture/api/http"
	domain "clean-architecture/domain"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Http         *http.Server
	Router       *gin.Engine
	configLoader domain.IConfig
	App          *domain.App
	controllers  map[string]domain.IController
}

func New() *Server {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	server := &Server{
		App: &domain.App{},
		Http: &http.Server{
			Handler: router,
		},
		Router:      router,
		controllers: make(map[string]domain.IController),
	}

	router.Use(gin.Recovery())

	return server
}

func (s *Server) WithLogger(logger domain.ILogger) *Server {
	s.App.Logger = logger
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
	if s.App.Config, err = s.configLoader.Load(); err != nil {
		return err
	}

	s.Http.Addr = fmt.Sprintf(":%d", s.App.Config.Http.Port)

	for _, c := range s.controllers {
		if c != nil {
			if err = c.Setup(s.App); err != nil {
				return err
			}
		}
	}

	return routes.Register(s.App, s.Router, s.controllers)
}

func (s *Server) Start() (err error) {
	if err = s.Setup(); err != nil {
		return err
	}

	return s.Http.ListenAndServe()
}

func (s *Server) Stop() (err error) {
	return s.Http.Shutdown(context.Background())
}
