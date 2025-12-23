package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/misafari/rlingo/internal/delivery/http/handler"
	"github.com/misafari/rlingo/internal/delivery/http/response"
	"github.com/misafari/rlingo/internal/domain/translation"
	usecase "github.com/misafari/rlingo/internal/usecase/translation"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct{}

func (m *mockRepo) CreateNewTranslation(ctx context.Context, tr *translation.Translation) error {
	return nil
}

func TestCreateTranslation_Integration(t *testing.T) {
	app := fiber.New()

	repo := &mockRepo{}
	uc := usecase.NewModifyingUseCase(repo)
	h := handler.NewTranslationHandler(uc)

	app.Post("/translations", h.CreateTranslation)

	t.Run("should return 201 on valid request", func(t *testing.T) {
		reqBody := map[string]string{
			"key":    "welcome_msg",
			"locale": "en",
			"text":   "Welcome!",
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/translations", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var result response.SuccessResponse
		err = json.NewDecoder(resp.Body).Decode(&result)

		assert.NoError(t, err)
		assert.Equal(t, "translation created successfully", result.Message)

		data := result.Data.(map[string]interface{})
		assert.NotNil(t, data["ID"])
		assert.Equal(t, "welcome_msg", data["Key"])
	})

	t.Run("should return 400 on key is required", func(t *testing.T) {
		reqBody := map[string]string{
			"locale": "en",
			"text":   "Welcome!",
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/translations", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var result response.ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&result)

		assert.NoError(t, err)
		assert.Equal(t, "field validation error", result.Error)
		assert.Equal(t, "field 'Key' failed on the 'required' tag", result.Message)
	})
}
