package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/domain"
	projectHandler "github.com/misafari/rlingo/internal/project/handler"
	"github.com/misafari/rlingo/internal/share/middleware"
	"github.com/misafari/rlingo/tests/project/mock"
)

func TestProjectHandler_GetAll_ReturnsOK(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	mockService := mock.NewMockProjectServiceWithFetchAllFn(
		[]*domain.Project{{ID: uuid.MustParse("ea772f3c-a5a0-4121-a714-981f40222ef4"), Name: "Core"}},
		nil,
	)

	h := projectHandler.NewHttpHandler(mockService)
	app.Get("/projects", h.GetAll)

	req := httptest.NewRequest(http.MethodGet, "/projects", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var got []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if len(got) != 1 {
		t.Fatalf("expected 1 project, got %d", len(got))
	}
	if got[0]["name"] != "Core" {
		t.Fatalf("expected project name %q, got %v", "Core", got[0]["name"])
	}
}
