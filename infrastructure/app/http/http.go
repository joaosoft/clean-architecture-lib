package http

import (
	routes "clean-architecture/api/http"
	controller "clean-architecture/controllers/http"
	"clean-architecture/domain"
	"clean-architecture/infrastructure/config"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHttp struct {
	http         *http.Server
	router       *gin.Engine
	logger       domain.ILogger
	db           *sql.DB
	configLoader domain.IConfig
	config       *config.Config
}

func New() *AppHttp {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	app := &AppHttp{
		http: &http.Server{
			Handler: router,
		},
		router: router,
	}

	router.Use(gin.Recovery())

	return app
}

func (s *AppHttp) Start() (err error) {
	if s.configLoader != nil {
		if s.config, err = s.configLoader.Load(); err != nil {
			return err
		}
	}
	s.http.Addr = fmt.Sprintf(":%d", s.config.Http.Port)
	return s.http.ListenAndServe()
}

func (s *AppHttp) Stop() (err error) {
	return s.http.Shutdown(context.Background())
}

func (s *AppHttp) WithController(controller ...controller.IController) domain.IApp {
	routes.RegisterRoutes(s, controller...)
	return s
}

func (s *AppHttp) WithLogger(logger domain.ILogger) domain.IApp {
	s.logger = logger
	return s
}

func (s *AppHttp) Logger() domain.ILogger {
	return s.logger
}

func (s *AppHttp) WithDb(db *sql.DB) domain.IApp {
	s.db = db
	return s
}

func (s *AppHttp) Db() *sql.DB {
	return s.db
}

func (s *AppHttp) WithHttp(http *http.Server) domain.IApp {
	s.http = http
	return s
}

func (s *AppHttp) Http() *http.Server {
	return s.http
}

func (s *AppHttp) WithRouter(router *gin.Engine) domain.IApp {
	s.router = router
	return s
}

func (s *AppHttp) Router() *gin.Engine {
	return s.router
}

func (s *AppHttp) WithConfig(config *config.Config) domain.IApp {
	s.config = config
	return s
}

func (s *AppHttp) WithConfigLoader(configLoader domain.IConfig) domain.IApp {
	s.configLoader = configLoader
	return s
}

func (s *AppHttp) Config() *config.Config {
	return s.config
}
