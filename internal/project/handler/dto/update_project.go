package dto

type UpdateProjectRequest struct {
	Name     string    `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=100"`
}

func (r *UpdateProjectRequest) ToEntity() (*domain.Project, error) {
	return &domain.Project{
		Name: r.Name,
		Description: r.Description,
	}, nil
}