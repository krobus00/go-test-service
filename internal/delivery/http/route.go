package http

import "github.com/labstack/echo/v4"

type HttpHandler struct {
	SimpleMathHandler *SimpleMathHandler
	Echo              *echo.Echo
}

func NewHttpHandler() *HttpHandler {
	return new(HttpHandler)
}

func (h *HttpHandler) RegisterSimpleMathHandler(handler *SimpleMathHandler) {
	h.SimpleMathHandler = handler
}

func (h *HttpHandler) RegisterEcho(e *echo.Echo) {
	h.Echo = e
}

func (h *HttpHandler) RegisterHttpRoute() {
	e := h.Echo

	e.POST("/math/add", h.SimpleMathHandler.HandleAdd)
	e.POST("/math/sub", h.SimpleMathHandler.HandleSub)
	e.POST("/math/div", h.SimpleMathHandler.HandleDiv)
	e.POST("/math/multi", h.SimpleMathHandler.HandleMulti)
}
