package enrollment

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/usecases/project_enroller"
)

func Routes(g *echo.Group) func(enroller project_enroller.ProjectEnroller) {
	return func(enroller project_enroller.ProjectEnroller) {
		g.POST("/enroll-to-project", enrollToProject(enroller))
	}
}

func enrollToProject(enroller project_enroller.ProjectEnroller) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request ApplicationRequest
		if err := c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		enrollmentApplication := fromRequestToApplication(&request)
		application, appErr := enroller.EnrollToProject(&enrollmentApplication)
		if appErr != nil {
			return appErr
		}

		response := fromApplicationToResponse(application)

		return c.JSON(http.StatusAccepted, response)
	}
}
