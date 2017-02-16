package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	log "github.com/Sirupsen/logrus"
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

func logExecCommand(command string, planet *Planet, strucOut *StructuredOuput) {
	log.Debugln("### execCommand complete ###")
	log.Debugf("user: %s\n", planet.user)
	log.Debugf("hostname: %s\n", planet.host)
	log.Debugf("orbit key: %s\n", os.Getenv("ORBIT_KEY"))
	log.Debugf("command: %s\n", command)
	log.Debugf("strucOut: %v\n", strucOut)
	log.Debugf("planet: %s\n maxLineLength: %d\n", strucOut.planet, strucOut.maxOutLength)
}
