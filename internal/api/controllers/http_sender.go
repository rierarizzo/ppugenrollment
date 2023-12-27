package controllers

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/pkg/domain"
)

type errorResponse struct {
	StatusCode int    `json:"status_code"`
	ErrorType  string `json:"error_type"`
	ErrorMsg   string `json:"error_msg"`
}

func sendError(statusCode int, appErr *domain.AppError) error {
	payload := errorResponse{
		StatusCode: statusCode,
		ErrorType:  appErr.Type,
		ErrorMsg:   appErr.Error(),
	}

	return echo.NewHTTPError(statusCode, payload)
}

type okResponse struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
	Data       any    `json:"data,omitempty"`
}

func sendOK(c echo.Context, statusCode int, msg string, response any) error {
	if msg == "" {
		msg = "Ok"
	}

	payload := okResponse{
		StatusCode: statusCode,
		Msg:        msg,
		Data:       response,
	}

	return c.JSON(statusCode, payload)
}
