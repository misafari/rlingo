package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/delivery/http/request"
	"github.com/misafari/rlingo/internal/project"
	"github.com/misafari/rlingo/internal/project/const"
	"github.com/misafari/rlingo/internal/project/domain"
	dto "github.com/misafari/rlingo/internal/project/handler/http_dto"
	"github.com/misafari/rlingo/internal/share"
)

type HttpHandler struct {
	service   *project.Service
	validator *validator.Validate
}

func (h *HttpHandler) GetAll(c *fiber.Ctx) error {
	all, err := h.service.FetchAll(c.Context())
	if err != nil {
		return share.InternalServerErrorResponse(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewProjectsResponseFromEntities(all))
}

func (h *HttpHandler) GetOneByID(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return share.ValidationErrorResponse(c, _const.ErrMissingProjectID)
	}

	projectUUID, err := uuid.Parse(id)
	if err != nil {
		return share.ValidationErrorResponse(c, err)
	}

	fetchedProject, err := h.service.FetchOneByID(c.Context(), projectUUID)
	if err != nil {
		return share.InternalServerErrorResponse(c, err)
	}

	if fetchedProject == nil {
		return share.NotFoundErrorResponse(c, "project not found")
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewProjectResponseFromEntity(fetchedProject))
}

func (h *HttpHandler) Create(c *fiber.Ctx) error {
	var req dto.UpdateProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return share.JsonParsingErrorResponse(c)
	}

	if err := h.validator.Struct(req); err != nil {
		return share.ValidationErrorResponse(c, err)
	}

	entity := &domain.Project{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.service.Create(c.Context(), entity); err != nil {
		return share.InternalServerErrorResponse(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.NewCreateProjectResponseFromEntity(entity))
}

func (h *HttpHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(share.ErrorResponse{
			Error: "bad_request", Message: "id is required",
		})
	}

	projectUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(share.ErrorResponse{
			Error: "bad_request", Message: "invalid id",
		})
	}

	var req request.SaveProjectRequest
	if err = c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(share.ErrorResponse{
			Error: "bad_request", Message: "cannot parse JSON",
		})
	}

	if err = h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(share.ErrorResponse{
			Error: "bad_request", Message: err.Error(),
		})
	}

	entity := &domain.Project{
		ID:   projectUUID,
		Name: req.Name,
	}

	err = h.service.Update(c.Context(), entity)
	if err != nil {
		return share.InternalServerErrorResponse(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewProjectResponseFromEntity(entity))
}

func (h *HttpHandler) DeleteOneById(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return share.ValidationErrorResponse(c, _const.ErrMissingProjectID)
	}

	projectUUID, err := uuid.Parse(id)
	if err != nil {
		return share.ValidationErrorResponse(c, err)
	}

	if err := h.service.DeleteOneById(c.Context(), projectUUID); err != nil {
		return share.InternalServerErrorResponse(c, err)
	}

	return c.Status(fiber.StatusNoContent).JSON(share.SuccessResponse{
		Message: "Project deleted successfully",
	})
}
