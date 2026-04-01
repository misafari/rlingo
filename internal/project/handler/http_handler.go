package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/dto"
	"github.com/misafari/rlingo/internal/project/service"
	error2 "github.com/misafari/rlingo/internal/project/error"
	domain "github.com/misafari/rlingo/internal/project/domain"
)

type HttpHandler struct {
	service *service.ProjectService
}

func (h *HttpHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return JsonParsingErrorResponse(c)
	}

	if err := h.validator.Struct(req); err != nil {
		return ValidationErrorResponse(c, err)
	}

	entity := &domain.Project{
		Name: req.Name,
		Description: req.Description,
	}
	
	if err := h.service.Create(c.Context(), entity); err != nil {
		return InternalServerErrorResponse(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.NewCreateProjectResponseFromEntity(entity))
}

func (h *ProjectHandler) DeleteOneById(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return ValidationErrorResponse(c, error2.ErrMissingProjectID)
	}

	projectUUID, err := uuid.Parse(id)
	if err != nil {
		return ValidationErrorResponse(c, err)
	}

	if err := h.service.DeleteOneById(c.Context(), projectUUID); err != nil {
		return InternalServerErrorResponse(c, err)
	}

	return c.Status(fiber.StatusNoContent).JSON(SuccessResponse{
		Message: "Project deleted successfully",
	})
}

func (h *ProjectHandler) FetchAll(c *fiber.Ctx) error {
	all, err := h.service.FetchAll(c.Context())
	if err != nil {
		return InternalServerErrorResponse(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewProjectsResponseFromEntity(all))
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
		return InternalServerErrorResponse(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewProjectResponseFromEntity(entity))
}