package errors

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   any    `json:"value,omitempty"`
}

type ValidationError struct {
	Fields []FieldError `json:"fields"`
}

func (v *ValidationError) Error() string {
	return "validation failed"
}

func NewValidationError(fields []FieldError) *ValidationError {
	return &ValidationError{Fields: fields}
}
