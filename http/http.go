package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/joaosoft/clean-infrastructure/domain/message"

	"github.com/opentracing/opentracing-go"

	"github.com/opentracing-contrib/go-gin/ginhttp"

	"github.com/gin-gonic/gin"
	"github.com/joaosoft/clean-infrastructure/config"
	"github.com/joaosoft/clean-infrastructure/domain"
	"github.com/joaosoft/clean-infrastructure/errors"
	httpConfig "github.com/joaosoft/clean-infrastructure/http/config"
)

// Http service
type Http struct {
	// Name
	name string
	// App
	app domain.IApp
	// Configurations
	config *httpConfig.Config
	// Http
	http *http.Server
	// Router
	router *gin.Engine
	// Routes
	routes []func() error
	// Controllers
	controllers []domain.IController
	// Middlewares
	middlewares []domain.IMiddleware
	// Recovery
	recovery gin.HandlerFunc
	// Additional Config Type
	additionalConfigType interface{}
	// Status Channel
	statusChannel chan error
	// Started
	started bool
}

const (
	// configFile http config file
	configFile = "http.yaml"
)

// New creates a new http service
func New(app domain.IApp, config *httpConfig.Config) *Http {
	engine := gin.New()
	newHttp := &Http{
		name: "Http",
		app:  app,
		http: &http.Server{
			Handler: engine,
		},
		statusChannel: make(chan error),
	}

	newHttp.WithRouter(engine)

	if config != nil {
		newHttp.config = config
	}

	return newHttp
}

// Name gets the service name
func (h *Http) Name() string {
	return fmt.Sprintf("%s server ready: %d", h.name, h.config.Port)
}

// Start starts the http service
func (h *Http) Start() (err error) {
	if h.config == nil {
		h.config = &httpConfig.Config{}
		h.config.AdditionalConfig = h.additionalConfigType
		if err = config.Load(h.ConfigFile(), h.config); err != nil {
			err = errors.ErrorLoadingConfigFile().Formats(h.ConfigFile(), err)
			message.ErrorMessage(h.Name(), err)
			return err
		}
	}

	// recovery
	if h.recovery == nil {
		h.recovery = h.newDefaultRecovery()
	}
	h.router.Use(h.recovery)

	// load body to context
	h.router.Use(loadBody)

	// prometheus metrics
	h.router.GET("/metrics", prometheusHandler())

	h.router.Use(ginhttp.Middleware(opentracing.GlobalTracer()))

	// register middlewares
	for _, middleware := range h.middlewares {
		middleware.RegisterMiddlewares()
	}

	// register routes
	for _, controller := range h.controllers {
		controller.Register()
	}

	if h.app.Config().Env == domain.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	for _, f := range h.routes {
		if err = f(); err != nil {
			return err
		}
	}

	h.http.Addr = fmt.Sprintf("%s:%d", h.config.Host, h.config.Port)

	go func(status chan error) {
		if err = h.http.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				status <- err
			}
		}
	}(h.statusChannel)

	h.started = true

	select {
	case err = <-h.statusChannel:
		return err
	case <-time.After(2 * time.Second):
		return nil
	}
}

// Config gets the http configurations
func (h *Http) Config() *httpConfig.Config {
	return h.config
}

// ConfigFile gets the config file
func (h *Http) ConfigFile() string {
	return configFile
}

// Stop stops the http server
func (h *Http) Stop() (err error) {
	if !h.started {
		return nil
	}
	defer close(h.statusChannel)
	if err = h.http.Shutdown(context.Background()); err != nil {
		return err
	}

	h.started = false
	return nil
}

// WithMiddleware adds a new controller to the server
func (h *Http) WithMiddleware(middleware domain.IMiddleware) domain.IHttp {
	h.middlewares = append(h.middlewares, middleware)
	return h
}

// WithController adds a new controller to the server
func (h *Http) WithController(controler domain.IController) domain.IHttp {
	h.controllers = append(h.controllers, controler)
	return h
}

// WithRouter sets the router engine
func (h *Http) WithRouter(router *gin.Engine) domain.IHttp {
	h.router = router
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, errors.ErrorUndefinedRoute())
	})
	return h
}

// Router gets the router engine
func (h *Http) Router() *gin.Engine {
	return h.router
}

// WithRecovery applies a recovery func
func (h *Http) WithRecovery(recovery gin.HandlerFunc) domain.IHttp {
	h.recovery = recovery
	return h
}

// WithAdditionalConfigType sets an additional config type
func (h *Http) WithAdditionalConfigType(obj interface{}) domain.IHttp {
	h.additionalConfigType = obj
	return h
}

// Started true if started
func (h *Http) Started() bool {
	return h.started
}
