package validator

import (
	"clean-architecture/controllers/http/validators"

	"github.com/joaosoft/validator"
)

func InitValidator() {
	validator.SetValidateAll(true).
		AddCallback("checkPerson", validators.CheckPerson)
}
