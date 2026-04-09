package dto

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}
