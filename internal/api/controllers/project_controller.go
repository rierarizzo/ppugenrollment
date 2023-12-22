package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/utils"
	"ppugenrollment/internal/ports"
)

type ProjectController struct {
	manager ports.ProjectManager
}

func NewProjectController(manager ports.ProjectManager) *ProjectController {
	return &ProjectController{manager}
}

func (pc *ProjectController) GetAllProjects(c echo.Context) error {
	projects, appErr := pc.manager.GetAllProjects()

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromProjectsToResponse(projects)

	return utils.SendOK(c, http.StatusOK, "", response)
}
