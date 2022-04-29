package domain

import (
	"clean-architecture/infrastructure/config"
)

type IConfig interface {
	Load() (_ *config.Config, err error)
}

type IController interface {
	Setup(app *App) error
}

type IModel interface {
	Setup(app *App) error
}

type IRepository interface {
	Setup(app *App) error
}

type IServer interface {
	Setup() error
	WithController(personController IController) IServer
	WithLogger(logger ILogger) IServer
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
