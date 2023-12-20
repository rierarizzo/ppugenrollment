package mappers

import (
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/domain"
)

func FromProjectToResponse(project *domain.Project) types.ProjectResponse {
	return types.ProjectResponse{
		ID:          project.ID,
		Company:     types.CompanyResponse{ID: project.Company.ID},
		Name:        project.Name,
		Description: project.Description,
		Starts:      project.Starts,
		Ends:        project.Ends,
	}
}

func FromProjectsToResponse(projects []domain.Project) []types.ProjectResponse {
	var response []types.ProjectResponse
	for _, v := range projects {
		response = append(response, FromProjectToResponse(&v))
	}

	return response
}
