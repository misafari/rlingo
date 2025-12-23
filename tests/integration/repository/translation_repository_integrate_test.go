package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/domain/translation"
	"github.com/misafari/rlingo/internal/repository/postgres"
	"github.com/stretchr/testify/assert"
)

func TestTranslateRepository_CreateNewTranslation(t *testing.T) {
	ctx := context.Background()
	repo := postgres.NewTranslateRepository(testPool)

	t.Run("insert success", func(t *testing.T) {
		id := uuid.New()
		tr := &translation.Translation{
			ID:     id,
			Key:    "hello",
			Locale: "en",
			Text:   "Hello World",
		}

		err := repo.CreateNewTranslation(ctx, tr)
		assert.NoError(t, err)

		var exists bool
		err = testPool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM translation WHERE id=$1)", id).Scan(&exists)
		assert.NoError(t, err)
		assert.True(t, exists)
	})
}
