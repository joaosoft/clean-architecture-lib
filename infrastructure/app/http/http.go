package http

import (
	"context"
	"database/sql"
	"fmt"
	routes "github.com/joaosoft/clean-architecture/api/http"
	"github.com/joaosoft/clean-architecture/infrastructure/config"
	"github.com/joaosoft/clean-architecture/infrastructure/domain/app"
	configDomain "github.com/joaosoft/clean-architecture/infrastructure/domain/config"
	httpDomain "github.com/joaosoft/clean-architecture/infrastructure/domain/http"
	"github.com/joaosoft/clean-architecture/infrastructure/domain/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHttp struct {
	http         *http.Server
	router       *gin.Engine
	logger       logger.ILogger
	db           *sql.DB
	configLoader configDomain.IConfig
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

func (s *AppHttp) WithController(controller ...httpDomain.IHttpController) app.IApp {
	routes.RegisterRoutes(s, controller...)
	return s
}

func (s *AppHttp) WithLogger(logger logger.ILogger) app.IApp {
	s.logger = logger
	return s
}

func (s *AppHttp) Logger() logger.ILogger {
	return s.logger
}

func (s *AppHttp) WithDb(db *sql.DB) app.IApp {
	s.db = db
	return s
}

func (s *AppHttp) Db() *sql.DB {
	return s.db
}

func (s *AppHttp) WithHttp(http *http.Server) app.IApp {
	s.http = http
	return s
}

func (s *AppHttp) Http() *http.Server {
	return s.http
}

func (s *AppHttp) WithRouter(router *gin.Engine) app.IApp {
	s.router = router
	return s
}

func (s *AppHttp) Router() *gin.Engine {
	return s.router
}

func (s *AppHttp) WithConfig(config *config.Config) app.IApp {
	s.config = config
	return s
}

func (s *AppHttp) WithConfigLoader(configLoader configDomain.IConfig) app.IApp {
	s.configLoader = configLoader
	return s
}

func (s *AppHttp) Config() *config.Config {
	return s.config
}
