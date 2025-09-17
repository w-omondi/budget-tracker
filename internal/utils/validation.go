package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseAndValidateData(ctx *fiber.Ctx, data any) error {
	if err := ctx.BodyParser(data); err != nil {
		return err
	}

	var _errors []string
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			_errors = append(_errors, err.Field())
		}
		return errors.New("Invalid: " + strings.Join(_errors, ", "))
	}

	return nil
}
