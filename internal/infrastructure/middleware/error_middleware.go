package middleware

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		statusCode := fiber.StatusInternalServerError
		var message any

		if e, ok := err.(*fiber.Error); ok {
			statusCode = e.Code
			message = e.Error()
		}

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make([]string, 0)
			for _, value := range validationErrors {
				errorMessages = append(errorMessages, fmt.Sprintf(
					"[%s]: '%v' | needs to implements '%s'",
					value.Field(),
					value.Value(),
					value.ActualTag(),
				))
			}

			statusCode = fiber.StatusBadRequest
			message = errorMessages
		}

		return c.Status(statusCode).JSON(fiber.Map{
			"errors": message,
		})
	}
}
