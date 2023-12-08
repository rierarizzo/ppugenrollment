package approval

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/usecases/enrollment_application_approver"
	"strconv"
)

func Routes(g *echo.Group) func(approver enrollment_application_approver.Approver) {
	return func(approver enrollment_application_approver.Approver) {
		g.POST("/approveEnrollmentApplication/:application_id", approveEnrollmentApplication(approver))
	}
}

func approveEnrollmentApplication(approver enrollment_application_approver.Approver) echo.HandlerFunc {
	return func(c echo.Context) error {
		applicationID, err := strconv.Atoi(c.Param("application_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		generated, appErr := approver.ApproveEnrollmentApplication(applicationID, 1)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}

		response := fromGeneratedToResponse(generated)

		return c.JSON(http.StatusAccepted, response)
	}
}
