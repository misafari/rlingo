package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project"
	"github.com/misafari/rlingo/internal/project/domain"
	projectHandler "github.com/misafari/rlingo/internal/project/handler"
	dto "github.com/misafari/rlingo/internal/project/handler/http_dto"
	appdto "github.com/misafari/rlingo/internal/share/dto"
	"github.com/misafari/rlingo/internal/share/middleware"
	"github.com/misafari/rlingo/tests/project/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateNewProject(t *testing.T) {
	tests := []struct {
		name                 string
		service              project.Service
		requestBody          string
		expectedResponse     *domain.Project
		expectedResponseCode int
		expectedErrorMessage string
	}{
		{
			"returns created when project is created",
			utils.NewMockProjectServiceWithCreateFn(&domain.Project{
				ID:          uuid.MustParse("4f7af2c7-6b95-4721-9cc5-c880027f34e5"),
				Name:        "Sample Project",
				Description: "This is a sample project description",
			}, nil),
			`{
				"name": "Sample Project",
				"description": "This is a sample project description"
			}`,
			&domain.Project{
				ID:          uuid.MustParse("4f7af2c7-6b95-4721-9cc5-c880027f34e5"),
				Name:        "Sample Project",
				Description: "This is a sample project description",
			},
			fiber.StatusCreated,
			"",
		},
		{
			"returns bad request when name is invalid (empty string)",
			nil,
			`{
				"name": "",
				"description": "This is a sample project description"
			}`,
			nil,
			fiber.StatusUnprocessableEntity,
			"Validation failed",
		},
		{
			"returns bad request when name is invalid (not present)",
			nil,
			`{
				"description": "This is a sample project description"
			}`,
			nil,
			fiber.StatusUnprocessableEntity,
			"Validation failed",
		},
		{
			"returns bad request when description is invalid (empty string)",
			nil,
			`{
				"name": "Sample Project",
				"description": ""
			}`,
			nil,
			fiber.StatusUnprocessableEntity,
			"Validation failed",
		},
		{
			"returns bad request when description is invalid (not present)",
			nil,
			`{
				"name": "Sample Project"
			}`,
			nil,
			fiber.StatusUnprocessableEntity,
			"Validation failed",
		},
		{
			"returns bad request when request body is invalid (empty body)",
			nil,
			"",
			nil,
			fiber.StatusBadRequest,
			"Invalid request body",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
			h := projectHandler.NewHttpHandler(tt.service)
			app.Post("/projects", h.Create)

			var body io.Reader
			if tt.requestBody != "" {
				body = strings.NewReader(tt.requestBody)
			}

			req := httptest.NewRequest(http.MethodPost, "/projects", body)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			require.NoError(t, err)

			assert.Equal(t, tt.expectedResponseCode, resp.StatusCode)

			if tt.expectedResponseCode == fiber.StatusCreated {
				assert.NotNil(t, tt.expectedResponse)

				var got dto.CreateProjectResponse
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))

				assert.NotNil(t, got)

				assert.Equal(t, tt.expectedResponse.ID, got.ID)
				assert.Equal(t, tt.expectedResponse.Name, got.Name)
				assert.Equal(t, tt.expectedResponse.Description, got.Description)
			} else {
				var got appdto.ErrorResponse
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))

				assert.False(t, got.Success)
				assert.Equal(t, tt.expectedErrorMessage, got.Message)
			}
		})
	}
}
