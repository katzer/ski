package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
	rotor "github.com/lestrrat/go-file-rotatelogs"
	hook "github.com/rifflock/lfshook"
)

func setupLogger(customLogfile string, verbose bool) {
	level := log.InfoLevel // default level
	if verbose {
		level = log.DebugLevel
	}

	logDir := path.Join(os.Getenv("ORBIT_HOME"), "logs")
	createLogDirIfNecessary(logDir)
	logFile := path.Join(logDir, "ski.log") // default log file
	if len(customLogfile) > 0 {
		filename := filepath.Base(customLogfile)
		logFile = path.Join(logDir, filename)
	}

	formatter := getDefaultFormatter()

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

func getDefaultFormatter() log.Formatter {
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
}

func createLogDirIfNecessary(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0775|os.ModeDir); err != nil {
			// can't do anything
			os.Stderr.WriteString(fmt.Sprintf("%v", err))
		}
	}
}
