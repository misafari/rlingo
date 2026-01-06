package usecase

import (
	"context"

	"github.com/google/uuid"
)

type CreateNewTranslationLocaleInput struct {
	LocaleID uuid.UUID
	value    string
}

type CreateNewTranslationInput struct {
	ProjectId uuid.UUID
	Key       uuid.UUID
	Locales   []CreateNewTranslationLocaleInput
}

type TranslationCreationUseCase struct {
}

func (t *TranslationCreationUseCase) Execute(ctx context.Context, input *CreateNewTranslationInput) {

}

func NewTranslationCreationUseCase() *TranslationCreationUseCase {
	return &TranslationCreationUseCase{}
}
