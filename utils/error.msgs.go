package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MsgForTag(fe validator.FieldError) string {
	// fe.Field() -- for field name
	// fe.Param() -- for field param (tag for error)
	// fe.Value() -- for field value
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return fmt.Sprintf("%v is NOT a valid Email.", fe.Value())
	case "gt":
		return "Must be greater than " + fe.Param()
	}
	return fe.Error() // default error message
}

func FormValidations(err error) map[string]string {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make(map[string]string)
		for _, fe := range ve {
			out[fe.Field()] = MsgForTag(fe)
		}
		return out
	}
	return nil
}
