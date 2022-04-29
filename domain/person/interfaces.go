package person

import (
	"context"

	"github.com/gin-gonic/gin"
)

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

type IPersonController interface {
	GetPersonByID(ctx *gin.Context)
}

type IPersonModel interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}

type IPersonRepository interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}
