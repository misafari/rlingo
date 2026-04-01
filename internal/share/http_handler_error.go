package share

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

var (
	ErrInvalidJson      = ErrorResponse{
		Error:   "bad_request",
		Message: "cannot parse JSON",
	}
)

func JsonParsingErrorResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(ErrInvalidJson)
}

func ValidationErrorResponse(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Error:   "bad_request",
		Message: err.Error(),
	})
}

func InternalServerErrorResponse(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
		Error:   "internal_server_error",
		Message: err.Error(),
	})
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
