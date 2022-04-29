package person

import (
	"context"

	"github.com/gin-gonic/gin"
)

type IPersonController interface {
	GetPersonByID(ctx *gin.Context)
}

type IPersonModel interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}

type IPersonRepository interface {
	GetPersonByID(ctx context.Context, personID int) (*Person, error)
}
