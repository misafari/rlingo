package request

type SaveLocaleRequest struct {
	ProjectID string `json:"project_id" validate:"required"`
	Locale    string `json:"locale"  validate:"required,min=3,max=100"`
	IsDefault bool   `json:"is_default" default:"false"`
}
