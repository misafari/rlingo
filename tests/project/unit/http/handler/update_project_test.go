package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	projectHandler "github.com/misafari/rlingo/internal/project/handler"
	appdto "github.com/misafari/rlingo/internal/share/dto"
	"github.com/misafari/rlingo/internal/share/middleware"
)

func TestProjectHandler_Update_ReturnsBadRequestWhenIDMissing(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	h := &projectHandler.HttpHandler{}

	app.Put("/projects/:id?", h.Update)

	req := httptest.NewRequest(http.MethodPut, "/projects", nil)
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
