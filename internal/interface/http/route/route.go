package route

import (
	"github.com/givxl33t/go-fiber-boilerplate/internal/interface/http/handler"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, userHandler *handler.UserHandler, authMiddleware fiber.Handler) {
	prefixRouter := app.Group("/api/v1")

	prefixRouter.Post("/users", userHandler.Register)
	prefixRouter.Post("/users/login", userHandler.Login)
	prefixRouter.Get("/users/current", authMiddleware, userHandler.Current)
	prefixRouter.Patch("/users/current", authMiddleware, userHandler.Update)
}
