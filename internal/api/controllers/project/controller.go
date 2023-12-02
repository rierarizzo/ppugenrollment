package project

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/usecases/project_manager"
)

func Routes(g *echo.Group) func(mngr project_manager.DefaultManager) {
	return func(mngr project_manager.DefaultManager) {
		g.POST("/get-all-projects", getAllProjects(mngr))
	}
}

func getAllProjects(mngr project_manager.DefaultManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		projects, appErr := mngr.GetAllProjects()
		if appErr != nil {
			return appErr
		}

		return c.JSON(http.StatusAccepted, fromProjectsToResponse(projects))
	}
}
