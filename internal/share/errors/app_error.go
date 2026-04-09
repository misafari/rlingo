package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFound(msg string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: msg}
}

func NotFoundF(format string, a ...any) *AppError {
	msg := fmt.Sprintf(format, a...)
	return &AppError{Code: http.StatusNotFound, Message: msg}
}

func BadRequest(msg string, details interface{}) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: msg, Details: details}
}

func Unauthorized(msg string) *AppError {
	return &AppError{Code: http.StatusUnauthorized, Message: msg}
}

func Forbidden(msg string) *AppError {
	return &AppError{Code: http.StatusForbidden, Message: msg}
}

func Internal(msg string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: msg}
}

func Conflict(msg string) *AppError {
	return &AppError{Code: http.StatusConflict, Message: msg}
}
