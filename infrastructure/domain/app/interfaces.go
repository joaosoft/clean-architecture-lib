package app

import (
	configImpl "clean-architecture/infrastructure/config"
	"clean-architecture/infrastructure/domain/config"
	httpController "clean-architecture/infrastructure/domain/http"
	"clean-architecture/infrastructure/domain/logger"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IApp interface {
	WithConfigLoader(configLoader config.IConfig) IApp
	WithConfig(config *configImpl.Config) IApp
	Config() *configImpl.Config
	WithLogger(logger logger.ILogger) IApp
	Logger() logger.ILogger
	WithDb(db *sql.DB) IApp
	Db() *sql.DB
	WithHttp(http *http.Server) IApp
	Http() *http.Server
	WithRouter(router *gin.Engine) IApp
	Router() *gin.Engine
	WithController(controller ...httpController.IHttpController) IApp
	Start() error
	Stop() error
}
