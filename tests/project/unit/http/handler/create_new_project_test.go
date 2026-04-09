package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/domain"
	projectHandler "github.com/misafari/rlingo/internal/project/handler"
	dto "github.com/misafari/rlingo/internal/project/handler/http_dto"
	appdto "github.com/misafari/rlingo/internal/share/dto"
	"github.com/misafari/rlingo/internal/share/middleware"
	"github.com/misafari/rlingo/tests/project/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectHandler_Create_ReturnsBadRequestWhenRequestBodyInvalid(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	h := &projectHandler.HttpHandler{}

	app.Post("/projects", h.Create)

	req := httptest.NewRequest(http.MethodPost, "/projects", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var got appdto.ErrorResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))

	assert.False(t, got.Success)
	assert.Equal(t, "Invalid request body", got.Message)
}

func TestProjectHandler_Create_ReturnsBadRequestWhenNameIsEmpty(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	h := projectHandler.NewHttpHandler(nil)

	app.Post("/projects", h.Create)

	reqBody := `{
		"name": "",
		"description": "This is a sample project description"
	}`

	req := httptest.NewRequest(http.MethodPost, "/projects", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)

	var got appdto.ErrorResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))

	assert.False(t, got.Success)
	assert.Equal(t, "Validation failed", got.Message)
}

func TestProjectHandler_Create_ReturnsBadRequestWhenDescriptionInvalid(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	h := &projectHandler.HttpHandler{}

	app.Post("/projects", h.Create)

	reqBody := `{
		"name": "Sample Project",
		"description": ""
	}`

	req := httptest.NewRequest(http.MethodPost, "/projects", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)

	var got appdto.ErrorResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))

	assert.False(t, got.Success)
	assert.Equal(t, "Validation failed", got.Message)
}

func TestProjectHandler_Create_ReturnsCreatedWhenProjectCreated(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})

	projectID := uuid.MustParse("4f7af2c7-6b95-4721-9cc5-c880027f34e5")

	mockService := mock.NewMockProjectServiceWithCreateFn(&domain.Project{
		ID:          projectID,
		Name:        "Sample Project",
		Description: "This is a sample project description",
	}, nil)

	h := projectHandler.NewHttpHandler(mockService)

	app.Post("/projects", h.Create)

	reqBody := `{
		"name": "Sample Project",
		"description": "This is a sample project description"
	}`

	req := httptest.NewRequest(http.MethodPost, "/projects", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var got dto.CreateProjectResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))

	assert.Equal(t, projectID, got.ID)
	assert.Equal(t, "Sample Project", got.Name)
	assert.Equal(t, "This is a sample project description", got.Description)
}
