package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misafari/rlingo/internal/identity"
	"github.com/misafari/rlingo/internal/identity/domain"
	dto "github.com/misafari/rlingo/internal/identity/handler/http_dto"
	apperror "github.com/misafari/rlingo/internal/share/errors"
	"github.com/misafari/rlingo/internal/share/utils"
)

type HttpHandler struct {
	service identity.UserService
}

func (h *HttpHandler) Signup(c *fiber.Ctx) error {
	var req dto.UserSignupRequest
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	entity := &domain.User{
		Email:        req.Email,
		PasswordHash: req.Password,
		FullName:     req.FullName,
		Status:       domain.UsersStatusACTIVE,
	}

	savedProject, err := h.service.SignUp(c.Context(), entity)
	if err != nil {
		return apperror.Internal(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(
		dto.NewUserSignupResponseFromEntity(savedProject),
	)
}

func NewUserHttpHandler(service identity.UserService) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}
