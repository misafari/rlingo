package request

import (
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/delivery/http/response"
	"github.com/misafari/rlingo/internal/domain/translation"
)

type SaveTranslationRequest struct {
	KeyID    string `json:"key_id" validate:"required"`
	LocaleID string `json:"locale_id" validate:"required"`
	Text     string `json:"text" validate:"required"`
}

func (s *SaveTranslationRequest) ToEntity() (*translation.Translation, *response.ErrorResponse) {
	keyID, err := uuid.Parse(s.KeyID)
	if err != nil {
		return nil, &response.ErrorResponse{
			Error: "bad_request", Message: "invalid key id",
		}
	}

	localeID, err := uuid.Parse(s.LocaleID)
	if err != nil {
		return nil, &response.ErrorResponse{
			Error: "bad_request", Message: "invalid locale id",
		}
	}

	return &translation.Translation{
		KeyID:    keyID,
		LocaleID: localeID,
		Text:     s.Text,
	}, nil
}
