package integration

import (
	"testing"

	"github.com/givxl33t/go-fiber-boilerplate/config"
	"github.com/givxl33t/go-fiber-boilerplate/internal/domain"
	"github.com/givxl33t/go-fiber-boilerplate/internal/infrastructure"
	"github.com/givxl33t/go-fiber-boilerplate/internal/infrastructure/middleware"
	"github.com/givxl33t/go-fiber-boilerplate/internal/interface/http/handler"
	"github.com/givxl33t/go-fiber-boilerplate/internal/interface/http/route"
	"github.com/givxl33t/go-fiber-boilerplate/internal/repository"
	"github.com/givxl33t/go-fiber-boilerplate/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type e2eTestSuite struct {
	suite.Suite
	Config         *viper.Viper
	App            *fiber.App
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository repository.UserRepository
	UserUsecase    usecase.UserUsecase
	UserHandler    *handler.UserHandler
	AuthMiddleware fiber.Handler
}

func TestE2eSuite(t *testing.T) {
	suite.Run(t, new(e2eTestSuite))
}

func (s *e2eTestSuite) SetupSuite() {
	s.Config = config.New()
	s.DB = infrastructure.NewGorm(s.Config)
	s.Log = infrastructure.NewLogger(s.Config)
	s.App = infrastructure.NewFiber(s.Config)
	s.Validate = infrastructure.NewValidator(s.Config)
	s.UserRepository = repository.NewUserRepository(s.DB)
	s.UserUsecase = usecase.NewUserUsecase(s.UserRepository, s.Log, s.Validate, s.Config)
	s.UserHandler = handler.NewUserHandler(s.UserUsecase, s.Log)
	s.AuthMiddleware = middleware.NewAuth(s.UserUsecase, s.Log)
	route.RegisterRoute(s.App, s.UserHandler, s.AuthMiddleware)
}

func (s *e2eTestSuite) SetupTest() {
	s.Require().NoError(s.DB.Migrator().AutoMigrate(&domain.User{}))
}

func (s *e2eTestSuite) TearDownTest() {
	s.Require().NoError(s.DB.Migrator().DropTable("users"))
}
