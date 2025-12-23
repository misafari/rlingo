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
	modifyingUseCase *usecase.TranslationModifyingUseCase
	validator        *validator.Validate
}

func NewTranslationHandler(modifyingUseCase *usecase.TranslationModifyingUseCase) *TranslationHandler {
	return &TranslationHandler{
		modifyingUseCase: modifyingUseCase,
		validator:        validator.New(),
	}
}

func (h *TranslationHandler) CreateTranslation(c *fiber.Ctx) error {
	var req request.CreateTranslationRequest
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
		ID:     uuid.New(),
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

	return c.Status(fiber.StatusCreated).JSON(response.SuccessResponse{
		Message: "translation created successfully",
		Data:    tr,
	})
}
