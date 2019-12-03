package api

import (
	"github.com/danijel-bjelancevic/golang-echo-api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//NewRouter will initialise routes and middlewares and return a new echo instance
func (h Handler) NewRouter() *echo.Echo {
	e := echo.New()

	h.defineMiddlewares(e)
	h.defineRoutes(e)

	return e
}

func (h Handler) defineRoutes(e *echo.Echo) {
	e.GET("/admin/ping", h.Livetest)
	e.GET("/admin/healthcheck", h.Healthcheck)
}

func (h Handler) defineMiddlewares(e *echo.Echo) {
	e.Use(mdw.NewTransactionIdMiddleware)
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
}
