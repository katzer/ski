package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

// StructuredOuput ...
type StructuredOuput struct {
	planet       string
	output       string
	maxOutLength int
}

// Default logfile path
const logFile string = "testlog.log"

func main() {
	args := os.Args
	opts := Opts{}
	opts.procArgs(args)

	if opts.helpFlag {
		printUsage()
		os.Exit(0)
	}
	if opts.versionFlag {
		printVersion()
		os.Exit(0)
	}

	level := log.InfoLevel
	if opts.debugFlag {
		level = log.DebugLevel
	}

	file, err := setupLogger(logFile, level)
	// No error -> a file was opened for writing, arange for closing the file
	if err == nil {
		defer closeFile(file)
	}

	log.Infof("Started with args: %v\n", os.Args)
	log.Debugln(&opts)
	exec := makeExecutor(&opts)
	exec.execMain(&opts)
	log.Infof("Ended with args: %v\n", os.Args)
}

func closeFile(file *os.File) {
	file.Sync()
	file.Close()
}

func makeExecutor(opts *Opts) Executor {
	executor := Executor{}
	var planet Planet
	for _, entry := range opts.planets {
		var user, host, dbID string
		connDet := getPlanetDetails(entry)
		planet.outputStruct.planet = entry
		id := entry
		planetType := getType(entry)
		user, host = getUserAndHost(connDet)
		switch planetType {
		case linuxServer:
			dbID = ""
		case database:
			dbID, connDet = procDBDets(connDet)
		case webServer:
			log.Warnln("Usage of goo with web servers is not implemented")
			continue
		default:
			log.Warnln("Unkown Type of target %s: %s\n", entry, planet.planetType)
			continue
		}
		executor.planets = append(executor.planets, Planet{id, user, host, planetType, dbID, StructuredOuput{id, "", 0}})
	}
	return executor
}

/**
*	Prints the current Version of the goo application
 */
func printVersion() {
	// TODO: Read it from a config file
	os.Stdout.WriteString("0.9\n")
}

/**
*	Prints the help dialog
 */
func printUsage() {
	usage := `usage: goo [options...] <planets>...
	Options:
	-s="<scriptname>"   Execute script and return result
	-c="<command>"      Execute script and return result
	-t=<"templatename>" Templatefile to be applied
	-p    Pretty print output as a table
	-l    Load bash profiles on Server
	-t    Show type of planet
	-h    Display this help text
	-v    Show version number
	-d    Show extended debug informations, set logging level to debug
`
	fmt.Println(usage)
}
