package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//Healthcheck will return 200 if all external connections are working
func (h Handler) Healthcheck(ctx echo.Context) error {
	h.Logger.Info("app is healthy")

	return ctx.JSON(http.StatusOK, "healthy")
}
