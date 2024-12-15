package exception

import "github.com/gofiber/fiber/v2"

// define error here
var (
	// user domain error shouldnt be here but its tightly coupled with fiber as of now
	ErrUserNotFound         = fiber.NewError(fiber.StatusNotFound, "user is not found")
	ErrUserAlreadyExist     = fiber.NewError(fiber.StatusBadRequest, "username already exist")
	ErrUserPasswordMismatch = fiber.NewError(fiber.StatusBadRequest, "password do not match")
	ErrUserUnauthorized     = fiber.NewError(fiber.StatusUnauthorized, "user unauthorized")

	// generic error
	ErrInternalServerError = fiber.ErrInternalServerError
)
