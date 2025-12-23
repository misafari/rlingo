package request

type CreateTranslationRequest struct {
	Key    string `json:"key" validate:"required"`
	Locale string `json:"locale" validate:"required"`
	Text   string `json:"text" validate:"required"`
}
