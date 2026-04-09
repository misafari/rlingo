package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project"
	"github.com/misafari/rlingo/internal/project/domain"
	dto "github.com/misafari/rlingo/internal/project/handler/http_dto"
	appdto "github.com/misafari/rlingo/internal/share/dto"
	apperror "github.com/misafari/rlingo/internal/share/errors"
	"github.com/misafari/rlingo/internal/share/utils"
)

type HttpHandler struct {
	service project.Service
}

func NewHttpHandler(service project.Service) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (h *HttpHandler) GetAll(c *fiber.Ctx) error {
	all, err := h.service.FetchAll(c.Context())
	if err != nil {
		return apperror.Internal(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewProjectsResponseFromEntities(all))
}

func (h *HttpHandler) GetOneByID(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return apperror.BadRequest("id is request", nil)
	}

	projectUUID, err := uuid.Parse(id)
	if err != nil {
		return apperror.BadRequest("uuid is not valid", err)
	}

	fetchedProject, err := h.service.FetchOneByID(c.Context(), projectUUID)
	if err != nil {
		return apperror.Internal(err.Error())
	}

	if fetchedProject == nil {
		return apperror.NotFoundF("project not found with id: %s", id)
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewProjectResponseFromEntity(fetchedProject))
}

func (h *HttpHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateProjectRequest
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	entity := &domain.Project{
		Name:        req.Name,
		Description: req.Description,
	}

	savedProject, err := h.service.Create(c.Context(), entity)
	if err != nil {
		return apperror.Internal(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.NewCreateProjectResponseFromEntity(savedProject))
}

func (h *HttpHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return apperror.BadRequest("id is request", nil)
	}

	projectUUID, err := uuid.Parse(id)
	if err != nil {
		return apperror.BadRequest("id is not valid", nil)
	}

	var req dto.CreateProjectRequest
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	entity := &domain.Project{
		ID:   projectUUID,
		Name: req.Name,
	}

	err = h.service.Update(c.Context(), entity)
	if err != nil {
		return apperror.Internal(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(dto.NewProjectResponseFromEntity(entity))
}

func (h *HttpHandler) DeleteOneById(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return apperror.BadRequest("id is request", nil)
	}

	projectUUID, err := uuid.Parse(id)
	if err != nil {
		return apperror.BadRequest("id is not valid", nil)
	}

	if err = h.service.DeleteOneById(c.Context(), projectUUID); err != nil {
		return apperror.Internal(err.Error())
	}

	return c.Status(fiber.StatusNoContent).JSON(appdto.SuccessResponse{
		Message: "Project is deleted successfully",
	})
}
