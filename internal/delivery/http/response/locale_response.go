package response

import "github.com/misafari/rlingo/internal/domain/locale"

type LocaleResponse struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Locale    string `json:"locale"`
	IsDefault bool   `json:"is_default"`
}

func NewLocaleResponseFromEntity(entity *locale.Locale) LocaleResponse {
	return LocaleResponse{
		ID:        entity.ID.String(),
		ProjectID: entity.ProjectID.String(),
		Locale:    entity.Locale.String(),
		IsDefault: entity.IsDefault,
	}
}

func NewLocaleResponseListFromEntity(entityList []*locale.Locale) []LocaleResponse {
	l := len(entityList)
	if l < 1 {
		return []LocaleResponse{}
	}

	response := make([]LocaleResponse, l)
	for index, le := range entityList {
		if le == nil {
			continue
		}
		response[index] = NewLocaleResponseFromEntity(le)
	}

	return response
}
