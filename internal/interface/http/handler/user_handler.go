package handler

import (
	"github.com/givxl33t/go-fiber-boilerplate/internal/model"
	"github.com/givxl33t/go-fiber-boilerplate/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
	Logger      *logrus.Logger
}

func NewUserHandler(userUsecase usecase.UserUsecase, log *logrus.Logger) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
		Logger:      log,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Registers a new user with the given details.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body model.RegisterUserRequest true "User Registration Request"
// @Success 201 {object} model.UserResponse
// @Router /users [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	registerUserRequest := new(model.RegisterUserRequest)
	if err := c.BodyParser(registerUserRequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	response, err := h.UserUsecase.Register(c.Context(), registerUserRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error user register")
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(&model.WebResponse[*model.UserResponse]{
			Data: response,
		})
}

// Login godoc
// @Summary User Login
// @Description Logs in a user and returns a JWT token.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body model.LoginUserRequest true "User Login Request"
// @Success 200 {object} model.TokenResponse
// @Router /users/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	loginUserRequest := new(model.LoginUserRequest)
	if err := c.BodyParser(loginUserRequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	response, err := h.UserUsecase.Login(c.Context(), loginUserRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error user login")
		return err
	}

	return c.
		JSON(&model.WebResponse[*model.TokenResponse]{
			Data: response,
		})
}

// Current godoc
// @Summary Get Current User
// @Description Retrieves the currently logged-in user's information.
// @Tags Users
// @Produce json
// @Success 200 {object} model.UserResponse
// @Security BearerAuth
// @Router /users/current [get]
func (h *UserHandler) Current(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	response, err := h.UserUsecase.Current(c.Context(), &model.GetUserRequest{Username: auth.Username})
	if err != nil {
		h.Logger.WithError(err).Error("error update user")
		return err
	}

	return c.
		JSON(&model.WebResponse[*model.UserResponse]{
			Data: response,
		})
}

// Update godoc
// @Summary Update Current User
// @Description Updates the currently logged-in user's profile.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body model.UpdateUserRequest true "User Update Request"
// @Success 200 {object} model.UserResponse
// @Security BearerAuth
// @Router /users/current [patch]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	updateUserRequest := new(model.UpdateUserRequest)
	if err := c.BodyParser(updateUserRequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	updateUserRequest.Username = auth.Username
	response, err := h.UserUsecase.Update(c.Context(), updateUserRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error update user")
		return err
	}

	return c.
		JSON(&model.WebResponse[*model.UserResponse]{
			Data: response,
		})
}
