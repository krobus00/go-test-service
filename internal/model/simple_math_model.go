package model

import (
	"context"
	"errors"
)

var (
	ErrZeroNotAllowed = errors.New("zero not allowed")
)

type SimpleMathUsecase interface {
	Add(ctx context.Context, a float64, b float64) (float64, error)
	Sub(ctx context.Context, a float64, b float64) (float64, error)
	Div(ctx context.Context, a float64, b float64) (float64, error)
	Multi(ctx context.Context, a float64, b float64) (float64, error)
}

type (
	MathRequest struct {
		A float64 `json:"a"`
		B float64 `json:"b"`
	}
	MathResponse struct {
		Result float64 `json:"result"`
	}
)
