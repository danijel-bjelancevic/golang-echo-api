package main

import (
	"fmt"
	"github.com/danijel-bjelancevic/golang-echo-api/api"
	"github.com/danijel-bjelancevic/golang-echo-api/config"
	"github.com/danijel-bjelancevic/golang-echo-api/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		logrus.WithError(err).Fatal("can not get configurations")
	}

	logr, err := logger.New(cfg.Debug, cfg.ApiName, cfg.GrayLogUDPAddr)
	if err != nil {
		logrus.WithError(err).Fatal("can not initialize logger")
	}

	// Api dependencies
	h := &api.Handler{
		AppName: cfg.ApiName,
		Logger:  logr,
	}

	// start the server, and log if it fails
	e := h.NewRouter()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ApiPort)))
}
