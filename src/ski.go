package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
)

func parseOptions() Opts {
	var help, pretty, debug, load, version bool
	var jobFile, template, scriptName, command string

	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&pretty, "p", false, "prettyprint")
	flag.BoolVar(&debug, "d", false, "verbose")
	flag.BoolVar(&load, "l", false, "ssh profile loading")
	flag.BoolVar(&version, "v", false, "version")
	flag.StringVar(&jobFile, "j", "", "path to a json file with a task description")
	flag.StringVar(&template, "t", "", "filename of template")
	flag.StringVar(&scriptName, "s", "", "name of the script(regardless wether db or bash) to be executed")
	flag.StringVar(&command, "c", "", "command to be executed in quotes")
	flag.Parse()

	if len(jobFile) > 0 {
		return createATaskFromJobFile(jobFile)
	}
	opts := Opts{
		Help:       help,
		Pretty:     pretty,
		Debug:      debug,
		Load:       load,
		Version:    version,
		Template:   template,
		ScriptName: scriptName,
		Command:    command,
		Planets:    flag.Args(),
	}

	return opts
}

func main() {
	opts := parseOptions()
	validate(&opts)
	opts.postProcessing()

	level := log.InfoLevel
	if opts.Debug {
		level = log.DebugLevel
	}

	// Default logfile path
	logDir := path.Join(os.Getenv("ORBIT_HOME"), "logs")
	createLogDirIfNecessary(logDir)
	logFile := path.Join(logDir, "ski.log")
	setupLogger(logFile, level)

	log.Infof("Started with args: %v", os.Args)
	log.Debug(&opts)
	exec := makeExecutor(&opts)
	exec.execMain(&opts)
	log.Infof("Ended with args: %v", os.Args)
}

func makeExecutor(opts *Opts) Executor {
	log.Debugf("Function: makeExecutor")
	executor := Executor{}
	for _, planetID := range opts.Planets {
		planet := parseConnectionDetails(planetID)
		if !isValidPlanet(planet) {
			os.Exit(1) // TODO ask if it really is wanted.
		}
		planet.id = planetID
		planet.outputStruct = StructuredOuput{planetID, "", 0}
		executor.planets = append(executor.planets, planet)
	}
	log.Debugf("executor: %s", executor)
	return executor
}

func isValidPlanet(planet Planet) bool {
	ok := isSupported(planet.planetType)
	if !ok {
		switch planet.planetType {
		case webServer:
			os.Stderr.WriteString("Usage of ski with web servers is not implemented")
		default:
			os.Stderr.WriteString("Unkown Type of target")
		}
	}
	// TODO: since we know what kind of action is attempted on this server
	// we could check if the action is permitted on the current planet and
	// if not mark it as not valid
	return ok
}

/**
*	Prints the help dialog
 */
func printUsage() {
	usage := `usage: ski [options...] <planets>...
	Options:
	-s="<scriptname>"   Execute script and return result
	-c="<command>"      Execute script and return result
	-t=<"templatename>" Templatefile to be applied
	-p    Pretty print output as a table
	-l    Load bash profiles on Server
	-h    Display this help text
	-v    Show version number
	-d    Show extended debug informations, set logging level to debug
`
	fmt.Println(usage)
}
