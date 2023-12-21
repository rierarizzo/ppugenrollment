package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/domain"
	auth "ppugenrollment/internal/ports"
)

func AuthRoutes(g *echo.Group) func(userAuth auth.Authenticator) {
	return func(userAuth auth.Authenticator) {
		g.POST("/register", register(userAuth))
		g.POST("/login", login(userAuth))
	}
}

// register is a function that registers a new user
func register(userAuth auth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request types.UserRegisterRequest
		if err := c.Bind(&request); err != nil {
			slog.Error(err.Error())
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		if err := validateStruct(request); err != nil {
			slog.Error(err.Error())
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		user := mappers.FromRegisterRequestToUser(&request)
		appErr := userAuth.Register(&user)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}
		return c.JSON(http.StatusAccepted, "OK")
	}
}

// login is a function that logs in a user
func login(userAuth auth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request types.UserLoginRequest
		if err := c.Bind(&request); err != nil {
			slog.Error(err.Error())
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		if err := validateStruct(request); err != nil {
			slog.Error(err.Error())
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		authPayload, appErr := userAuth.Login(request.Email, request.Password)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}
		return c.JSON(http.StatusAccepted, mappers.FromAuthPayloadToResponse(authPayload))
	}
}

// validateStruct is a function that validates a struct using a new validator
func validateStruct(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
