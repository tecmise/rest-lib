package exceptions

import (
	"errors"
	"fmt"
	"time"
)

type ErrorType int

const (
	// Tipo indefinido
	TypeUnknown ErrorType = iota
	TypeInternal
	TypeValidation
	TypeBusiness
	TypeNotFound
)

type AppError struct {
	Type        ErrorType
	Code        ErrorCode
	Message     string
	Time        time.Time
	OriginalErr error
}

// Implementa a interface 'error'
func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.OriginalErr
}

func NewInternal(code ErrorCode, message string, err error) *AppError {
	if message == "" {
		message = "Ocorreu um erro interno no servidor."
	}
	return &AppError{
		Code:        code,
		Type:        TypeInternal,
		Message:     message,
		OriginalErr: err,
	}
}

func NewValidation(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Type:    TypeValidation,
		Message: message,
	}
}

func NewBusiness(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Type:    TypeBusiness,
		Message: message,
	}
}

func NewNotFound(code ErrorCode, resourceName string) *AppError {
	return &AppError{
		Type:    TypeNotFound,
		Message: fmt.Sprintf("%s não encontrado(a).", resourceName),
		Time:    time.Now(),
	}
}

func GetType(err error) ErrorType {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Type
	}
	return TypeUnknown
}
