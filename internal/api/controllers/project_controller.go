package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/ports"
)

func ProjectRoutes(g *echo.Group) func(mngr ports.Manager) {
	return func(mngr ports.Manager) {
		g.GET("/getAllProjects", getAllProjects(mngr))
	}
}

func getAllProjects(mngr ports.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		projects, appErr := mngr.GetAllProjects()

		if appErr != nil {
			return sendError(http.StatusInternalServerError, appErr)
		}

		response := mappers.FromProjectsToResponse(projects)

		return sendOK(c, http.StatusOK, "", response)
	}
}
