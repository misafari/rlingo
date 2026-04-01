package share

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
