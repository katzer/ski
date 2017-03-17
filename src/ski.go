package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
)

var help, pretty, debug, load, _version, saveReport bool
var jobFile, logFile, template, scriptName, command string

func main() {
	opts := parseOptions()
	validateArgsCount(&opts)
	opts.validate()

	verbose := opts.Debug || len(opts.LogFile) > 0
	setupLogger(opts.LogFile, verbose)
	setupDirs()

	log.Infof("Started with args: %v", os.Args)
	log.Debug(&opts)
	exec := makeExecutor(&opts)
	exec.execMain(&opts)
	if len(jobFile) == 0 {
		formatAndPrint(exec.planets, &opts, os.Stdout)
		handleExitCode(exec.planets)
		return
	}
	options := map[string]string{}
	options["job_name"] = path.Base(jobFile)
	options["orbit_home"] = os.Getenv("ORBIT_HOME")
	options["output"] = "jobs_output"

	createJSONReport(options, exec.planets, &opts)
}

// if there were any errors during the execution of command or scripts
// on any of the planets given as parameter, the exit code is set to non zero
// in case other programs rely on the exit code.
func handleExitCode(planets []Planet) {
	for _, entry := range planets {
		structuralError := entry.outputStruct == nil
		executionFailure := entry.outputStruct.errored
		if structuralError || executionFailure {
			os.Exit(1)
		}
	}
}

func makeExecutor(opts *Opts) Executor {
	log.Debugf("Function: makeExecutor")
	executor := Executor{}
	planets := parseConnectionDetails(opts.Planets)
	executor.planets = planets
	log.Debugf("executor: %s", executor)
	return executor
}

func isValidPlanet(planet Planet) bool {
	ok := isSupported(planet.planetType)
	if !ok {
		switch planet.planetType {
		case webServer:
			fmt.Fprintf(os.Stderr, "Usage of ski with web servers is not implemented")
		default:
			fmt.Fprintf(os.Stderr, "Unkown Type of target")
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
	os.Exit(0)
}

func parseOptions() Opts {
	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&pretty, "p", false, "prettyprint")
	flag.BoolVar(&debug, "d", false, "verbose")
	flag.BoolVar(&load, "l", false, "ssh profile loading")
	flag.BoolVar(&_version, "v", false, "version")
	flag.BoolVar(&saveReport, "js", false, "if the summary should be saved in json format. Used with the job flag")
	flag.StringVar(&jobFile, "j", "", "path to a json file with a task description")
	flag.StringVar(&logFile, "logfile", "ski.log", "path to a file for logging")
	flag.StringVar(&template, "t", "", "filename of template")
	flag.StringVar(&scriptName, "s", "", "name of the script(regardless wether db or bash) to be executed")
	flag.StringVar(&command, "c", "", "command to be executed in quotes")
	// flag.Usage = printUsage
	flag.Parse()

	if len(jobFile) > 0 {
		retval := createATaskFromJobFile(jobFile)
		postProcessing(&retval)
		return retval
	}

	opts := Opts{
		Help:       help,
		Pretty:     pretty,
		Debug:      debug,
		Load:       load,
		Version:    _version,
		SaveReport: saveReport,
		Template:   template,
		ScriptName: scriptName,
		Command:    command,
		Planets:    flag.Args(),
		LogFile:    logFile,
	}

	postProcessing(&opts)

	return opts
}

func validateArgsCount(opts *Opts) {
	if opts.Version {
		printVersion()
		os.Exit(0)
	}
	tooFew := len(os.Args) == 1
	// TODO Check if flags package removes the leading and trailing white spaces.
	scriptEmpty := len(opts.ScriptName) == 0
	cmdEmpty := len(opts.Command) == 0
	if tooFew || scriptEmpty && cmdEmpty {
		printUsage()
	}
}

func postProcessing(opts *Opts) {
	opts.Command = strings.Trim(opts.Command, "\"")
	opts.Template = strings.Trim(opts.Template, "\"")
	opts.ScriptName = strings.Trim(opts.ScriptName, "\"")
}
func setupDirs() {
	makeDir("tmp")
	makeDir("chronJobs")
}
