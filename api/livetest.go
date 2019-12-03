package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//Livetest is there to check if the api is live
func (h Handler) Livetest(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "pong")
}
