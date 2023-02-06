package infrastructure

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/krobus00/go-test-service/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (i *Infrastructure) NewRouter() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.HTTPErrorHandler = i.customHttpHandler

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	e.GET("/readiness", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	i.Router = e
}

func (i *Infrastructure) customHttpHandler(err error, c echo.Context) {
	var (
		customErr *model.HttpCustomError
		echoErr   *echo.HTTPError
	)

	// default error as internal server error
	response := model.NewHttpCustomError(
		http.StatusInternalServerError,
		errors.New(http.StatusText(http.StatusInternalServerError)),
	)

	switch {
	case errors.As(err, &customErr):
		response.StatusCode = customErr.StatusCode
		if customErr.StatusCode != http.StatusInternalServerError {
			response.Message = customErr.Error()
		}
	case errors.As(err, &echoErr):
		if m, ok := echoErr.Message.([]model.ValidationError); ok {
			response.Message = "validation errors"
			response.Errors = m
			break
		}
		response.StatusCode = echoErr.Code
		if echoErr.Code != http.StatusInternalServerError {
			response.Message = fmt.Sprintf("%v", echoErr.Message)
		}
	}

	_ = c.JSON(response.StatusCode, response)
}
