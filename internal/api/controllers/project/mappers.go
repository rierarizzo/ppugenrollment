package project

import "ppugenrollment/internal/domain"

func fromProjectToResponse(project *domain.Project) Response {
	return Response{
		ID:          project.ID,
		Company:     project.Company,
		Description: project.Description,
		Schedule:    project.Schedule,
	}
}

func fromProjectsToResponse(projects []domain.Project) []Response {
	var response []Response
	for _, v := range projects {
		response = append(response, fromProjectToResponse(&v))
	}

	return response
}
