package validators

import (
	"errors"

	"github.com/joaosoft/validator"
)

func CheckPerson(context *validator.ValidatorContext, validationData *validator.ValidationData) []error {
	personID := validationData.Value.Int()
	if personID <= 0 {
		return []error{errors.New("the person ID is invalid!")}
	}

	return nil
}
