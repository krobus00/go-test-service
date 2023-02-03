package usecase

import (
	"context"

	"github.com/krobus00/go-test-service/internal/model"
)

type simpleMathUsecase struct{}

func NewSimpleMathUsecase() model.SimpleMathUsecase {
	return new(simpleMathUsecase)
}

func (u *simpleMathUsecase) Add(ctx context.Context, a float64, b float64) (float64, error) {
	return a + b, nil
}

func (u *simpleMathUsecase) Sub(ctx context.Context, a float64, b float64) (float64, error) {
	return a - b, nil
}

func (u *simpleMathUsecase) Div(ctx context.Context, a float64, b float64) (float64, error) {
	switch {
	case a == 0 && b == 0:
		return 0, model.ErrZeroNotAllowed
	case b == 0:
		return 0, model.ErrZeroNotAllowed
	default:
		return a / b, nil
	}
}

func (u *simpleMathUsecase) Multi(ctx context.Context, a float64, b float64) (float64, error) {
	return a * b, nil
}
