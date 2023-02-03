package http

import (
	"net/http"

	"github.com/krobus00/go-test-service/internal/model"
	"github.com/labstack/echo/v4"
)

type SimpleMathHandler struct {
	SimpleMathUsecase model.SimpleMathUsecase
}

func NewSimpleMathHandler() *SimpleMathHandler {
	return new(SimpleMathHandler)
}

func (h *SimpleMathHandler) RegisterSimpleMathUecase(u model.SimpleMathUsecase) {
	h.SimpleMathUsecase = u
}

func (h *SimpleMathHandler) HandleAdd(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()
	payload := new(model.MathRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	result, err := h.SimpleMathUsecase.Add(ctx, payload.A, payload.B)
	if err != nil {
		return err
	}

	resp := model.MathResponse{
		Result: result,
	}

	return eCtx.JSON(http.StatusOK, resp)
}

func (h *SimpleMathHandler) HandleSub(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()
	payload := new(model.MathRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	result, err := h.SimpleMathUsecase.Sub(ctx, payload.A, payload.B)
	if err != nil {
		return err
	}

	resp := model.MathResponse{
		Result: result,
	}

	return eCtx.JSON(http.StatusOK, resp)
}

func (h *SimpleMathHandler) HandleDiv(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()
	payload := new(model.MathRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	result, err := h.SimpleMathUsecase.Div(ctx, payload.A, payload.B)

	switch err {
	case nil:
	case model.ErrZeroNotAllowed:
		return model.NewHttpCustomError(http.StatusBadRequest, model.ErrZeroNotAllowed)
	default:
		return err
	}

	resp := model.MathResponse{
		Result: result,
	}

	return eCtx.JSON(http.StatusOK, resp)
}

func (h *SimpleMathHandler) HandleMulti(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()
	payload := new(model.MathRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	result, err := h.SimpleMathUsecase.Multi(ctx, payload.A, payload.B)
	if err != nil {
		return err
	}

	resp := model.MathResponse{
		Result: result,
	}

	return eCtx.JSON(http.StatusOK, resp)
}
