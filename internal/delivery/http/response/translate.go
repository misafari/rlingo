package response

import (
	"github.com/misafari/rlingo/internal/domain/translation"
)

type TranslateResponse struct {
	ID        string `json:"id"`
	Key       string `json:"key"`
	Locale    string `json:"locale"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewTranslateResponseFromEntity(entity *translation.Translation) *TranslateResponse {
	return &TranslateResponse{
		ID:        entity.ID.String(),
		Key:       entity.KeyID.String(),
		Locale:    entity.LocaleID.String(),
		Text:      entity.Text,
		CreatedAt: entity.CreatedAt.String(),
		UpdatedAt: entity.UpdatedAt.String(),
	}
}

func NewTranslatesResponseFromEntity(entity []*translation.Translation) []TranslateResponse {
	if len(entity) < 1 {
		return []TranslateResponse{}
	}

	var response = make([]TranslateResponse, len(entity))
	for index, t := range entity {
		response[index] = *NewTranslateResponseFromEntity(t)
	}

	return response
}
