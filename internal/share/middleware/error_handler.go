package middleware

import (
	"errors"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/misafari/rlingo/internal/share/dto"
	apperrors "github.com/misafari/rlingo/internal/share/errors"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return ctx.Status(fiberErr.Code).JSON(dto.ErrorResponse{
			Success: false,
			Message: fiberErr.Message,
		})
	}

	// 2. Handle validation errors from go-playground/validator
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		fields := make([]apperrors.FieldError, 0, len(validationErrs))
		for _, ve := range validationErrs {
			fields = append(fields, apperrors.FieldError{
				Field:   toSnakeCase(ve.Field()),
				Message: validationMessage(ve),
			})
		}
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Success: false,
			Message: "Validation failed",
			Details: fields,
		})
	}

	var valErr *apperrors.ValidationError
	if errors.As(err, &valErr) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Success: false,
			Message: "Validation failed",
			Details: valErr.Fields,
		})
	}

	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return ctx.Status(appErr.Code).JSON(dto.ErrorResponse{
			Success: false,
			Message: appErr.Message,
			Details: appErr.Details,
		})
	}

	// 5. Unknown errors — log and return generic 500
	slog.Error("unhandled error", "error", err, "path", ctx.Path())
	return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
		Success: false,
		Message: "Internal server error",
	})
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Must be a valid email address"
	case "min":
		return "Value is too short or too small (min: " + fe.Param() + ")"
	case "max":
		return "Value is too long or too large (max: " + fe.Param() + ")"
	case "len":
		return "Must be exactly " + fe.Param() + " characters"
	case "oneof":
		return "Must be one of: " + fe.Param()
	case "url":
		return "Must be a valid URL"
	case "uuid":
		return "Must be a valid UUID"
	case "numeric":
		return "Must be a numeric value"
	case "gte":
		return "Must be greater than or equal to " + fe.Param()
	case "lte":
		return "Must be less than or equal to " + fe.Param()
	default:
		return "Invalid value"
	}
}

func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, r+32)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
