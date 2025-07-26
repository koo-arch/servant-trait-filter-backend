package httperr

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Code string

const (
	CodeBadRequest       Code = "bad_request"
	CodeValidationFailed Code = "validation_failed"
	CodeUnauthorized     Code = "unauthorized"
	CodeForbidden        Code = "forbidden"
	CodeNotFound         Code = "not_found"
	CodeConflict         Code = "conflict"
	CodeTooManyRequests  Code = "too_many_requests"
	CodeInternal         Code = "internal"
	CodeUnavailable      Code = "service_unavailable"
)

type FieldError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type Response struct {
	Code        Code          `json:"code"`
	Message     string        `json:"message"`
	RequestID   string        `json:"requestId,omitempty"`
	FieldErrors []FieldError  `json:"fieldErrors,omitempty"`
	Details     any           `json:"details,omitempty"`
}

// アプリケーション用のラッパー
type AppError struct {
	Code       Code
	HTTPStatus int
	Message    string
	Details    any
	Err        error
}

func (e *AppError) Error() string { 
	if e.Err != nil { return e.Err.Error() }
	return e.Message
}
func (e *AppError) Unwrap() error { 
	return e.Err 
}

func New(code Code, status int, msg string) *AppError {
	return &AppError{
		Code:       code,
		HTTPStatus: status,
		Message:    msg,
	}
}
func Wrap(err error, code Code, status int, msg string) *AppError {
	return &AppError{
		Code:       code,
		HTTPStatus: status,
		Message:    msg,
		Err:       err,
	}
}

func Validation(message string, fields []FieldError) *AppError {
	return &AppError{
		Code:       CodeValidationFailed,
		HTTPStatus: http.StatusBadRequest,
		Message:    message,
		Details:    map[string]any{}, // 任意
		// FieldErrors は Response 側で付与
	}
}

// unknown を internal にマップ
func From(err error) *AppError {
	var app *AppError
	if errors.As(err, &app) {
		return app
	}
	return Wrap(err, CodeInternal, http.StatusInternalServerError, "内部エラーが発生しました。")
}

// レスポンス整形 + 送出
func Write(c *gin.Context, err error, opts ...func(*Response)) {
	app := From(err)

	reqID := c.GetString("requestId") // request-id ミドルウェアで設定想定
	resp := Response{
		Code:      app.Code,
		Message:   app.Message,
		RequestID: reqID,
	}

	// Validation の FieldErrors を details から取り出す
	if fe, ok := app.Details.([]FieldError); ok {
		resp.FieldErrors = fe
		resp.Details = nil
	} else {
		resp.Details = app.Details
	}

	for _, opt := range opts {
		opt(&resp)
	}

	c.AbortWithStatusJSON(app.HTTPStatus, resp)
}