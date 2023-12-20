package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
	"strconv"
)

func ApprovalRoutes(g *echo.Group) func(approver ports.Approver) {
	return func(approver ports.Approver) {
		g.POST("/approveEnrollmentApplication/:application_id", approveEnrollmentApplication(approver))
	}
}

// approveEnrollmentApplication processes the approval of an enrollment application.
// It takes an approver that has the capability to approve enrollment applications.
// It returns an echo.HandlerFunc which handles the approval request.
func approveEnrollmentApplication(approver ports.Approver) echo.HandlerFunc {
	return func(c echo.Context) error {
		applicationID, err := strconv.Atoi(c.Param("application_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		approvedBy := c.Get("UserID").(int)

		generated, appErr := approver.ApproveEnrollmentApplication(applicationID, approvedBy)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}

		response := mappers.FromGeneratedToResponse(generated)

		return c.JSON(http.StatusAccepted, response)
	}
}
