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

func EnrollmentRoutes(g *echo.Group) func(enroller ports.Enroller) {
	return func(enroller ports.Enroller) {
		g.POST("/enrollToProject", enrollToProject(enroller))
	}
}

// enrollToProject is a function that handles the enrollment of a student to a project. It takes an enroller and returns
// an echo.HandlerFunc that can be used as a handler for HTTP requests
func enrollToProject(enroller ports.Enroller) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request types.EnrollmentApplicationRequest

		if err := c.Bind(&request); err != nil {
			slog.Error(err.Error())
			return sendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
		}

		enrolledBy := c.Get("UserID").(int)

		if err := validateStruct(request); err != nil {
			slog.Error(err.Error())
			return sendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
		}

		enrollmentApplication := mappers.FromRequestToApplication(&request)
		application, appErr := enroller.EnrollToProject(&enrollmentApplication, enrolledBy)

		if appErr != nil {
			return sendError(http.StatusInternalServerError, appErr)
		}

		response := mappers.FromApplicationToResponse(application)

		return sendOK(c, http.StatusAccepted, "Enrollment applied", response)
	}
}
