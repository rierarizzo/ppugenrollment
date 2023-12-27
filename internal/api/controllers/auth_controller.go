package controllers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
)

type AuthController struct {
	userAuth ports.UserAuthenticator
}

func NewAuthController(userAuth ports.UserAuthenticator) *AuthController {
	return &AuthController{userAuth}
}

func (ac *AuthController) Register(c echo.Context) error {
	request := new(types.UserRegisterRequest)

	if err := c.Bind(&request); err != nil {
		slog.Error(err.Error())
		return sendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	if appErr := request.Validate(); appErr != nil {
		slog.Error(appErr.Error())
		return sendError(http.StatusBadRequest, appErr)
	}

	user := mappers.FromRegisterRequestToUser(request)
	appErr := ac.userAuth.Register(&user)

	if appErr != nil {
		return sendError(http.StatusInternalServerError, appErr)
	}

	return sendOK(c, http.StatusCreated, "User registered", nil)
}

func (ac *AuthController) Login(c echo.Context) error {
	request := new(types.UserLoginRequest)

	if err := c.Bind(&request); err != nil {
		slog.Error(err.Error())
		return sendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	if appErr := request.Validate(); appErr != nil {
		slog.Error(appErr.Error())
		return sendError(http.StatusBadRequest, appErr)
	}

	authPayload, appErr := ac.userAuth.Login(request.Email, request.Password)

	if appErr != nil {
		return sendError(http.StatusUnauthorized, appErr)
	}

	response := mappers.FromAuthPayloadToResponse(authPayload)

	return sendOK(c, http.StatusAccepted, "User logged", response)
}
