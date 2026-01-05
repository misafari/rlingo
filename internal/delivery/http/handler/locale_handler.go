package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/delivery/http/request"
	"github.com/misafari/rlingo/internal/delivery/http/response"
	"github.com/misafari/rlingo/internal/domain/locale"
	usecase "github.com/misafari/rlingo/internal/usecase/locale"
	"golang.org/x/text/language"
)

type LocaleHttpHandler struct {
	crudUseCase *usecase.CrudLocaleUseCase
	validator   *validator.Validate
}

func (h *LocaleHttpHandler) Create(c *fiber.Ctx) error {
	var req request.SaveLocaleRequest
	if err := c.BodyParser(&req); err != nil {
		return response.JsonParsingErrorResponse(c)
	}

	if err := h.validator.Struct(req); err != nil {
		return response.ValidationErrorResponse(c, err)
	}

	projectUUID, err := uuid.Parse(req.ProjectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "invalid project id",
		})
	}

	languageTag, err := language.Parse(req.Locale)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "invalid locale",
		})
	}

	entity := &locale.Locale{
		ProjectID: projectUUID,
		Locale:    languageTag,
	}

	err = h.crudUseCase.Create(c.Context(), entity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "internal_server_error", Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(req)
}

func (h *LocaleHttpHandler) FetchAll(c *fiber.Ctx) error {
	all, err := h.crudUseCase.FetchAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "internal_server_error", Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.NewLocaleResponseListFromEntity(all))
}

func (h *LocaleHttpHandler) DeleteOneById(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "id is required",
		})
	}

	localUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{Error: "bad_request", Message: "invalid id"})
	}

	return h.crudUseCase.DeleteOneById(c.Context(), localUUID)
}

func (h *LocaleHttpHandler) Update(c *fiber.Ctx) error {
	params := c.Params("id", "")
	if params == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "id is required",
		})
	}

	localUUID, err := uuid.Parse(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "invalid id",
		})
	}

	var req request.SaveLocaleRequest
	if err = c.BodyParser(&req); err != nil {
		return response.JsonParsingErrorResponse(c)
	}

	if err = h.validator.Struct(req); err != nil {
		return response.ValidationErrorResponse(c, err)
	}

	projectUUID, err := uuid.Parse(req.ProjectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "invalid project id",
		})
	}

	languageTag, err := language.Parse(req.Locale)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "invalid locale",
		})
	}

	entity := &locale.Locale{
		ID:        localUUID,
		ProjectID: projectUUID,
		Locale:    languageTag,
	}

	err = h.crudUseCase.Update(c.Context(), entity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "internal_server_error", Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(req)
}

func NewLocaleHttpHandler(crudUseCase *usecase.CrudLocaleUseCase) *LocaleHttpHandler {
	return &LocaleHttpHandler{
		crudUseCase: crudUseCase,
		validator:   validator.New(),
	}
}
