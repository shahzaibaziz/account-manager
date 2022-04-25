package client

import (
	"github.com/sirupsen/logrus"
)

func log() logrus.FieldLogger {
	return logrus.WithField("pkg", "client")
}
