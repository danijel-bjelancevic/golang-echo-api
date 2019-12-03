package api

import (
	"github.com/danijel-bjelancevic/golang-echo-api/responses"
	"github.com/sirupsen/logrus"
)

//Handler holds all dependencies and it's initialised in the main
type Handler struct {
	AppName string
	Logger  logrus.FieldLogger
}

func (h Handler) errorResponse(err error, code int, tid string, msg string, logMsg string) responses.Response {
	h.Logger.WithError(err).Error(logMsg)
	response := responses.BuildResponse(nil, code, msg, tid)
	return response
}
