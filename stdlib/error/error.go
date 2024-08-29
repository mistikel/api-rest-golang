package errors

import (
	"errors"
	"fmt"
)

const (
	CodeHTTPParamDecode = iota + 100
	CodeHTTPParamValidate
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("no records found")
)

type AppError struct {
	Code       int
	Message    string
	Err        error
	StatusCode int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, error: %v", e.Code, e.Message, e.Err)
}

func NewAppError(code int, message string, statusCode int, err error) *AppError {
	return &AppError{Code: code, Message: message, StatusCode: statusCode, Err: err}
}

func NewDatabaseError(message string, err error) *AppError {
	return NewAppError(500, message, 500, err)
}
