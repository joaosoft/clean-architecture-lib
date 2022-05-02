package domain

import (
	"clean-architecture/infrastructure/config"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IController interface {
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Post(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type IConfig interface {
	Load() (_ *config.Config, err error)
}

type IApp interface {
	WithConfigLoader(configLoader IConfig) IApp
	WithConfig(*config.Config) IApp
	Config() *config.Config
	WithLogger(logger ILogger) IApp
	Logger() ILogger
	WithDb(db *sql.DB) IApp
	Db() *sql.DB
	WithHttp(http *http.Server) IApp
	Http() *http.Server
	WithRouter(router *gin.Engine) IApp
	Router() *gin.Engine
	WithController(controller ...IController) IApp
	Start() error
	Stop() error
}

type ILogger interface {
	Printf(format string, v ...any)
	Print(v ...any)
	Println(v ...any)
	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)
	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
}
