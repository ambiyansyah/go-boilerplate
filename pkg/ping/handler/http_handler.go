package handler

import (
	"net/http"

	"github.com/ambiyansyah/go-boilerplate/pkg/config"
	"github.com/ambiyansyah/go-boilerplate/pkg/ping"
	"github.com/labstack/echo"
)

type httpHandler struct {
	config *config.AppConfig
	route  *echo.Group
}

// NewHTTPHandler ...
func NewHTTPHandler(config *config.AppConfig, route *echo.Group) ping.HandlerHTTPService {
	return &httpHandler{
		config: config,
		route:  route,
	}
}

func (h *httpHandler) Mount() {
	h.route.GET("/", h.Get)
}

func (h *httpHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
