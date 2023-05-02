package person

import (
	"clean-architecture/controllers/structs"
	"clean-architecture/controllers/structs/person"
	personDomain "clean-architecture/domain/person"
	appDomain "clean-architecture/infrastructure/domain/app"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joaosoft/validator"
)

type PersonController struct {
	app   appDomain.IApp
	model personDomain.IPersonModel
}

func NewPersonController(app appDomain.IApp, model personDomain.IPersonModel) personDomain.IPersonController {
	return &PersonController{
		app:   app,
		model: model,
	}
}

func (c *PersonController) Get(ctx *gin.Context) {
	fmt.Println("running person controller")

	ctx.Header("Content-Type", "application/json")

	personID, _ := strconv.Atoi(ctx.Param("id_person"))
	request := person.GetPersonByIDRequest{
		IdPerson: personID,
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		var errMessages []string
		for _, err := range errs {
			errMessages = append(errMessages, err.Error())
		}

		ctx.JSON(http.StatusBadRequest,
			structs.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: strings.Join(errMessages, ", "),
			})
		return
	}

	person, err := c.model.GetPersonByID(ctx.Request.Context(), request.IdPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			structs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, person)
}

func (c *PersonController) Put(ctx *gin.Context)    {}
func (c *PersonController) Post(ctx *gin.Context)   {}
func (c *PersonController) Delete(ctx *gin.Context) {}
