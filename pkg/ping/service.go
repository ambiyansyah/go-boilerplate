package ping

import "github.com/labstack/echo"

// HandlerHTTPService ...
type HandlerHTTPService interface {
	Mount()
	Get(c echo.Context) error
}
