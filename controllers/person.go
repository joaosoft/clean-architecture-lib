package controllers

import (
	"clean-architecture/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IController interface {
	GetPersonByID(ctx *gin.Context)
}

type Controller struct {
	model models.IModel
}

func NewController(model models.IModel) IController {
	return &Controller{
		model: model,
	}
}

func (c *Controller) GetPersonByID(ctx *gin.Context) {
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param("id_person"),
	}

	ctx.Header("Content-Type", "application/json")

	person, err := c.model.GetPersonByID(ctx, request.IdPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.JSON(http.StatusOK, person)
}
