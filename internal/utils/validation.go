package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct(_struct interface{}) []string {
	var errors []string

	err := validate.Struct(_struct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field())
		}
		return errors
	}

	return nil
}
