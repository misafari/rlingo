package http_dto

import "github.com/misafari/rlingo/internal/project/domain"

type ProjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UpdatedAt   string `json:"updated_at"`
}

func NewProjectResponseFromEntity(entity *domain.Project) *ProjectResponse {
	return &ProjectResponse{
		ID:          entity.ID.String(),
		Name:        entity.Name,
		Description: entity.Description,
		Status:      string(entity.Status),
		UpdatedAt:   entity.UpdatedAt.String(),
	}
}

func NewProjectsResponseFromEntities(entity []*domain.Project) []ProjectResponse {
	if len(entity) < 1 {
		return []ProjectResponse{}
	}

	var response = make([]ProjectResponse, len(entity))
	for index, p := range entity {
		response[index] = *NewProjectResponseFromEntity(p)
	}

	return response
}
