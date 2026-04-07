package http_dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/domain"
)

type CreateProjectRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=100"`
}

type CreateProjectResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCreateProjectResponseFromEntity(entity *domain.Project) *CreateProjectResponse {
	return &CreateProjectResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Status:      string(entity.Status),
		CreatedBy:   entity.CreatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
