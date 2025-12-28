package request

type SaveProjectRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
