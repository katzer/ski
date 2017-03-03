package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

func parseOptions() Opts {
	var help, pretty, debug, load, version, saveReport bool
	var jobFile, logFile, template, scriptName, command string

	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&pretty, "p", false, "prettyprint")
	flag.BoolVar(&debug, "d", false, "verbose")
	flag.BoolVar(&load, "l", false, "ssh profile loading")
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&saveReport, "js", false, "if the summary should be saved in json format. Used with the job flag")
	flag.StringVar(&jobFile, "j", "", "path to a json file with a task description")
	flag.StringVar(&logFile, "logfile", "ski.log", "path to a file for logging")
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
		SaveReport: saveReport,
		Template:   template,
		ScriptName: scriptName,
		Command:    command,
		Planets:    flag.Args(),
		LogFile:    logFile,
	}

	return opts
}

func main() {
	opts := parseOptions()
	opts.postProcessing()
	validate(&opts)

	if opts.Help {
		printUsage()
		os.Exit(0)
	}
	if opts.Version {
		printVersion()
		os.Exit(0)
	}

	verbose := opts.Debug || len(opts.LogFile) > 0
	setupLogger(opts.LogFile, verbose)

	log.Infof("Started with args: %v", os.Args)
	log.Debug(&opts)
	exec := makeExecutor(&opts)
	exec.execMain(&opts)
	log.Infof("Ended with args: %v", os.Args)
}

/**
func createLogDirIfNecessary(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0775|os.ModeDir); err != nil {
			// can't do anything
			os.Stderr.WriteString(fmt.Sprintf("%v", err))
		}
	}
}**/

func makeExecutor(opts *Opts) Executor {
	log.Debugf("Function: makeExecutor")
	executor := Executor{}
	for i, planetID := range opts.Planets {
		planet := parseConnectionDetails(planetID)
		if !isValidPlanet(planet) {
			os.Exit(1) // TODO ask if it really is wanted.
		}
		planet.id = planetID
		planet.outputStruct = &StructuredOuput{planetID, "", i}
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
