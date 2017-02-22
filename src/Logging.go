package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	rotor "github.com/lestrrat/go-file-rotatelogs"
	hook "github.com/rifflock/lfshook"
)

func setupLogger(customLogfile string, level log.Level) (*os.File, error) {
	logFile, err := os.OpenFile(customLogfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Warnf("logrus will write to stderr. Bad file name: %s", customLogfile)
	} else {
		log.SetOutput(logFile)
	}

	formatter := setupDefaultFormatter()
	log.SetFormatter(formatter)
	log.SetLevel(level)
	return logFile, err
}

func setupLoggerWithRotation(customLogfile string, level log.Level) {
	logFile := "ski.log" // default log file
	if customLogfile != "" {
		logFile = customLogfile
	}

	formatter := setupDefaultFormatter()

	log.SetFormatter(formatter)
	log.SetLevel(level)

	writer, err := rotor.New(
		logFile+".%Y%m%d%H%M",
		rotor.WithLinkName(logFile),
		rotor.WithMaxAge(24*time.Hour),
		rotor.WithRotationTime(time.Hour),
	)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v. Rolling file appender can't be used.\n", err))
		os.Stderr.WriteString("Logrus will log to stderr.")
	} else {
		hook := hook.NewHook(hook.WriterMap{
			log.DebugLevel: writer,
			log.WarnLevel:  writer,
			log.InfoLevel:  writer,
			log.ErrorLevel: writer,
			log.FatalLevel: writer,
		})
		log.AddHook(hook)
		log.SetOutput(ioutil.Discard)
	}
}

func setupLoggerWithFileAppender(customLogfile string, level log.Level) {
	logFile := "ski.log" // default log file
	if customLogfile != "" {
		logFile = customLogfile
	}

	formatter := setupDefaultFormatter()
	log.SetFormatter(formatter)
	log.SetLevel(level)

	hook := hook.NewHook(hook.PathMap{
		log.DebugLevel: logFile,
		log.WarnLevel:  logFile,
		log.InfoLevel:  logFile,
		log.ErrorLevel: logFile,
		log.FatalLevel: logFile,
	})

	log.AddHook(hook)
	log.SetOutput(ioutil.Discard)
}

func setupDefaultFormatter() log.Formatter {
	formatter := new(log.TextFormatter)
	formatter.DisableSorting = true
	formatter.ForceColors = true
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true

	return formatter
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
