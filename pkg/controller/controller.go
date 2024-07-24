package controller

import (
	"context"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Controller struct { //nolint:maligned
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func New(_ context.Context, param *Param) (*Controller, *Param, error) {
	if param.LogLevel != "" {
		lvl, err := logrus.ParseLevel(param.LogLevel)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"log_level": param.LogLevel,
			}).WithError(err).Error("the log level is invalid")
		}
		logrus.SetLevel(lvl)
	}

	return &Controller{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}, param, nil
}
