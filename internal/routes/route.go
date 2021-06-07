package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Code int
			Msg  string
		}{
			Code: 200,
			Msg:  "hello world",
		})
	})
}
