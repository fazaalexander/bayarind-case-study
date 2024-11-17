package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func Validation(request interface{}) error {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErrs {
				if e.Tag() == "required" && e.Field() == "Username" {
					message = fmt.Sprintf("Please enter your %s", e.Field())
				} else if e.Tag() == "required" {
					message = fmt.Sprintf("Please enter your %s", e.Field())
				} else if e.Tag() == "min" && e.Field() == "Password" {
					message = "passwords must be at least 8 characters long"
				}
			}
			return errors.New(message)
		}
	}
	return nil
}
