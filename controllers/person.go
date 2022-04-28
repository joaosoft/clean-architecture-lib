package controllers

import (
	"clean-architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	model domain.IModel
}

func NewController(model domain.IModel) domain.IController {
	return &Controller{
		model: model,
	}
}

func (c *Controller) GetPersonByID(ctx *gin.Context) {
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param("id_person"),
	}

	ctx.Header("Content-Type", "application/json")

	person, err := c.model.GetPersonByID(ctx.Request.Context(), request.IdPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.JSON(http.StatusOK, person)
}
