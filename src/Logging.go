package main

import (
	"os"

	"github.com/Sirupsen/logrus"
)

func setupLogger(destination string, level logrus.Level) (*os.File, error) {
	logFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Warnf("logrus will write to stderr. Bad file name: %s", destination)
	} else {
		logrus.SetOutput(logFile)
	}

	formatter := new(logrus.TextFormatter)
	formatter.DisableSorting = true
	formatter.ForceColors = true
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true
	logrus.SetFormatter(formatter)
	logrus.SetLevel(level)
	return logFile, err
}
