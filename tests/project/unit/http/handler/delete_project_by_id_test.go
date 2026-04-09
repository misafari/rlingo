package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	projectHandler "github.com/misafari/rlingo/internal/project/handler"
	appdto "github.com/misafari/rlingo/internal/share/dto"
	"github.com/misafari/rlingo/internal/share/middleware"
	"github.com/misafari/rlingo/tests/project/mock"
)

func TestProjectHandler_DeleteOneById_ReturnsBadRequestWhenIDMissing(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	h := &projectHandler.HttpHandler{}

	app.Delete("/projects/:id?", h.DeleteOneById)

	req := httptest.NewRequest(http.MethodDelete, "/projects", nil)
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

func TestProjectHandler_DeleteOneById_ReturnsNoContentWhenDeleted(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	projectID := uuid.MustParse("4f7af2c7-6b95-4721-9cc5-c880027f34e5")
	mockService := mock.NewMockProjectServiceWithDeleteOneByIdFn(nil)
	h := projectHandler.NewHttpHandler(mockService)

	app.Delete("/projects/:id", h.DeleteOneById)

	req := httptest.NewRequest(http.MethodDelete, "/projects/"+projectID.String(), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, resp.StatusCode)
	}
}
