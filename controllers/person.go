package controllers

import (
	"clean-architecture/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaosoft/validator"
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
	personID, _ := strconv.Atoi(ctx.Param("id_person"))
	request := GetPersonByIDRequest{
		IdPerson: personID,
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		ctx.JSON(http.StatusBadRequest,
			ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: errs[0].Error(),
			})
		return
	}

	ctx.Header("Content-Type", "application/json")

	person, err := c.model.GetPersonByID(ctx.Request.Context(), request.IdPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, person)
}
