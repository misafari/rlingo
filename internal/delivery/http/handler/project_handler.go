package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/delivery/http/request"
	"github.com/misafari/rlingo/internal/delivery/http/response"
	"github.com/misafari/rlingo/internal/domain/project"
	usecase "github.com/misafari/rlingo/internal/usecase/project"
)

type ProjectHandler struct {
	crudUseCase *usecase.CrudProjectUseCase
	validator   *validator.Validate
}

func (h *ProjectHandler) Create(c *fiber.Ctx) error {
	var req request.SaveProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "cannot parse JSON",
		})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: err.Error(),
		})
	}

	entity := &project.Project{
		Name: req.Name,
	}

	err := h.crudUseCase.Create(c.Context(), entity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "internal_server_error", Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.NewProjectResponseFromEntity(entity),
	)
}

func (h *ProjectHandler) FetchAll(c *fiber.Ctx) error {
	all, err := h.crudUseCase.FetchAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "internal_server_error", Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.NewProjectsResponseFromEntity(all))
}

func (h *ProjectHandler) DeleteOneById(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "id is required",
		})
	}

	projectUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "invalid id",
		})
	}

	return h.crudUseCase.DeleteOneById(c.Context(), projectUUID)
}

func (h *ProjectHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "id is required",
		})
	}

	projectUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "invalid id",
		})
	}

	var req request.SaveProjectRequest
	if err = c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "cannot parse JSON",
		})
	}

	if err = h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: err.Error(),
		})
	}

	entity := &project.Project{
		ID:   projectUUID,
		Name: req.Name,
	}

	err = h.crudUseCase.Update(c.Context(), entity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error:   "",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.NewProjectResponseFromEntity(entity))
}

func NewProjectHandler(crudUseCase *usecase.CrudProjectUseCase) *ProjectHandler {
	return &ProjectHandler{
		crudUseCase: crudUseCase,
		validator:   validator.New(),
	}
}
