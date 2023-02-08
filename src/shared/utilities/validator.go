package utilities

import (
	"github.com/juanfer2/persona_5/src/shared/utilities/validations"
)

func ValidatorStruct(s any) []error {
	return validations.ValidateStruct(s)
}
