package infrastructure

import (
	"fmt"
	"time"

	_ "github.com/givxl33t/go-fiber-boilerplate/docs"
	"github.com/givxl33t/go-fiber-boilerplate/internal/infrastructure/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.GetString("APP_NAME"),
		ErrorHandler: middleware.NewErrorHandler(),
		Prefork:      config.GetBool("APP_PREFORK"),
		WriteTimeout: config.GetDuration("APP_TIMEOUT") * time.Second,
		ReadTimeout:  config.GetDuration("APP_TIMEOUT") * time.Second,
	})

	app.Get("/swagger/*", swagger.HandlerDefault)
	fmt.Println("Swagger UI available at /swagger/")

	app.Use(recover.New())

	return app
}
