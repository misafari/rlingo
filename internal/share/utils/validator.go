package utils

import (
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	apperrors "github.com/misafari/rlingo/internal/share/errors"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	})
	return validate
}

func ParseAndValidate(ctx *fiber.Ctx, dst interface{}) error {
	if err := ctx.BodyParser(dst); err != nil {
		return apperrors.BadRequest("Invalid request body", nil)
	}
	if err := GetValidator().Struct(dst); err != nil {
		return err // Let the global handler parse validator.ValidationErrors
	}
	return nil
}
