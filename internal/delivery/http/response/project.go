package response

import "github.com/misafari/rlingo/internal/domain/project"

type ProjectResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewProjectResponseFromEntity(entity *project.Project) *ProjectResponse {
	return &ProjectResponse{
		ID:   entity.ID.String(),
		Name: entity.Name,
	}
}

func NewProjectsResponseFromEntity(entity []*project.Project) []ProjectResponse {
	if len(entity) < 1 {
		return []ProjectResponse{}
	}

	var response = make([]ProjectResponse, len(entity))
	for index, p := range entity {
		response[index] = *NewProjectResponseFromEntity(p)
	}

	return response
}
