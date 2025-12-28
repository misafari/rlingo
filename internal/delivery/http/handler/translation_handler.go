package handler

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/delivery/http/request"
	"github.com/misafari/rlingo/internal/delivery/http/response"
	"github.com/misafari/rlingo/internal/domain/translation"
	usecase "github.com/misafari/rlingo/internal/usecase/translation"
)

type TranslationHandler struct {
	modifyingUseCase *usecase.CudTranslationUseCase
	readUseCase      *usecase.ReadTranslationUseCase
	validator        *validator.Validate
}

func NewTranslationHandler(
	modifyingUseCase *usecase.CudTranslationUseCase,
	readUseCase *usecase.ReadTranslationUseCase,
) *TranslationHandler {
	return &TranslationHandler{
		modifyingUseCase: modifyingUseCase,
		readUseCase:      readUseCase,
		validator:        validator.New(),
	}
}

func (h *TranslationHandler) Create(c *fiber.Ctx) error {
	var req request.SaveTranslationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "bad_request", Message: "cannot parse JSON",
		})
	}

	if err := h.validator.Struct(req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			fieldErr := validationErrors[0]
			dynamicMsg := fmt.Sprintf("field '%s' failed on the '%s' tag", fieldErr.Field(), fieldErr.Tag())

			return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error:   "field validation error",
				Message: dynamicMsg,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error:   "unknown error",
			Message: fmt.Sprintf("unknown error: %s", err.Error()),
		})
	}

	tr := &translation.Translation{
		Key:    req.Key,
		Locale: req.Locale,
		Text:   req.Text,
	}

	if err := h.modifyingUseCase.Create(c.Context(), tr); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.NewTranslateResponseFromEntity(tr))
}

func (h *TranslationHandler) FetchAll(c *fiber.Ctx) error {
	translations, err := h.readUseCase.FetchAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(response.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	if translations == nil || len(translations) == 0 {
		return c.Status(404).JSON(response.ErrorResponse{
			Error:   "not found",
			Message: "no translations found",
		})
	}

	return c.Status(200).JSON(response.NewTranslatesResponseFromEntity(translations))
}

func (h *TranslationHandler) DeleteOneById(c *fiber.Ctx) error {
	translationId := c.Params("id", "")

	if translationId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error:   "bad request",
			Message: "translation id is required",
		})
	}

	translationUUID, err := uuid.Parse(translationId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error:   "bad request",
			Message: "invalid translation id",
		})
	}

	err = h.modifyingUseCase.DeleteOneById(c.Context(), translationUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "translation deleted successfully",
		Data:    nil,
	})
}

func (h *TranslationHandler) Update(ctx *fiber.Ctx) error {
	translationId := ctx.Params("id", "")
	if translationId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error:   "bad request",
			Message: "translation id is required",
		})
	}

	translationUUID, err := uuid.Parse(translationId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error:   "bad request",
			Message: "invalid translation id",
		})
	}

	var req request.SaveTranslationRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error:   "bad_request",
			Message: "cannot parse JSON",
		})
	}

	if err := h.validator.Struct(req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			fieldErr := validationErrors[0]
			dynamicMsg := fmt.Sprintf("field '%s' failed on the '%s' tag", fieldErr.Field(), fieldErr.Tag())

			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error:   "field validation error",
				Message: dynamicMsg,
			})
		}
	}

	tr := &translation.Translation{
		ID:     translationUUID,
		Key:    req.Key,
		Locale: req.Locale,
	}

	if err = h.modifyingUseCase.Update(ctx.Context(), tr); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NewTranslateResponseFromEntity(tr))
}
