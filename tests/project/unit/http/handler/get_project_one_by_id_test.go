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
	appdto "github.com/misafari/rlingo/internal/share/dto"
	"github.com/misafari/rlingo/internal/share/middleware"
	"github.com/misafari/rlingo/tests/project/mock"
)

func TestProjectHandler_GetOneByID_ReturnsBadRequestWhenIDMissing(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	h := &projectHandler.HttpHandler{}

	app.Get("/projects/:id?", h.GetOneByID)

	req := httptest.NewRequest(http.MethodGet, "/projects", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var got appdto.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got.Success {
		t.Fatalf("expected success %v, got %v", false, got.Success)
	}

	if got.Message != "id is request" {
		t.Fatalf("expected message %q, got %q", "id is request", got.Message)
	}
}

func TestProjectHandler_GetOneByID_ReturnsBadRequestWhenIDInvalid(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	h := &projectHandler.HttpHandler{}

	app.Get("/projects/:id", h.GetOneByID)

	req := httptest.NewRequest(http.MethodGet, "/projects/not-a-uuid", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var got appdto.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got.Success {
		t.Fatalf("expected success %v, got %v", false, got.Success)
	}
	if got.Message != "uuid is not valid" {
		t.Fatalf("expected message %q, got %q", "uuid is not valid", got.Message)
	}
}

func TestProjectHandler_GetOneByID_ReturnsOKWhenProjectFound(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	projectID := uuid.MustParse("ef4ec8e9-f6bb-4ce7-bec6-fd32cec9afca")
	mockService := mock.NewMockProjectServiceWithFetchOneByIDFn(
		&domain.Project{
			ID:   projectID,
			Name: "Docs",
		},
		nil,
	)

	h := projectHandler.NewHttpHandler(mockService)

	app.Get("/projects/:id", h.GetOneByID)

	req := httptest.NewRequest(http.MethodGet, "/projects/"+projectID.String(), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var got map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got["id"] != projectID.String() {
		t.Fatalf("expected id %q, got %v", projectID.String(), got["id"])
	}
	if got["name"] != "Docs" {
		t.Fatalf("expected name %q, got %v", "Docs", got["name"])
	}
}
