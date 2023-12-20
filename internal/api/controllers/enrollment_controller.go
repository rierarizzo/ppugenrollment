package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
)

func EnrollmentRoutes(g *echo.Group) func(enroller ports.Enroller) {
	return func(enroller ports.Enroller) {
		g.POST("/enrollToProject", enrollToProject(enroller))
	}
}

func enrollToProject(enroller ports.Enroller) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request types.EnrollmentApplicationRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		enrolledBy := c.Get("UserID").(int)

		enrollmentApplication := mappers.FromRequestToApplication(&request)
		application, appErr := enroller.EnrollToProject(&enrollmentApplication, enrolledBy)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}

		response := mappers.FromApplicationToResponse(application)

		return c.JSON(http.StatusAccepted, response)
	}
}
