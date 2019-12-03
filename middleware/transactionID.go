package mdw

import (
	"github.com/labstack/echo/v4"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

const RequestId = "X-Request-id"

func NewTransactionIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware(next)
}

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tid := c.Request().Header.Get(RequestId)
		if tid == "" {
			tid = uuid.NewV4().String()
			c.Request().Header.Set(RequestId, tid)
		}
		c.Response().Header().Set(RequestId, tid)

		return next(c)
	}
}

func GetTransactionIdForLogger(c echo.Context) logrus.Fields {
	tid := c.Request().Header.Get(RequestId)
	c.Set("transaction_id", tid)
	return logrus.Fields{
		"transaction_id": tid,
	}
}

func GetTransactionId(c echo.Context) string {
	return c.Request().Header.Get(RequestId)
}
