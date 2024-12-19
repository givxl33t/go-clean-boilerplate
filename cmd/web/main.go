package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/givxl33t/go-fiber-boilerplate/config"
	"github.com/givxl33t/go-fiber-boilerplate/internal/infrastructure"
	"github.com/givxl33t/go-fiber-boilerplate/internal/infrastructure/middleware"
	"github.com/givxl33t/go-fiber-boilerplate/internal/interface/http/handler"
	"github.com/givxl33t/go-fiber-boilerplate/internal/interface/http/route"
	"github.com/givxl33t/go-fiber-boilerplate/internal/repository"
	"github.com/givxl33t/go-fiber-boilerplate/internal/usecase"
)

// @title Go Fiber Boilerplate API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config := config.New()

	app := infrastructure.NewFiber(config)
	port := config.Get("APP_PORT")

	db := infrastructure.NewGorm(config)
	logger := infrastructure.NewLogger(config)
	validate := infrastructure.NewValidator(config)
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, logger, validate, config)
	userHandler := handler.NewUserHandler(userUsecase, logger)

	authMiddleware := middleware.NewAuth(userUsecase, logger)

	route.RegisterRoute(app, userHandler, authMiddleware)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", port)); err != nil {
			panic(fmt.Errorf("error running app : %+v", err.Error()))
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch // blocks the main thread until an interrupt is received

	// cleanup tasks go here
	_ = app.Shutdown()

	fmt.Println("App shuts down successfully!")
}
