package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Opts ...
type Opts struct {
	debugFlag    bool
	helpFlag     bool
	loadFlag     bool
	prettyFlag   bool
	scriptFlag   bool
	tableFlag    bool
	typeFlag     bool
	versionFlag  bool
	planetsCount int
	command      string
	currentDBDet string
	currentDet   string
	scriptName   string
	template     string
	planets      []string
}

/**
*	Returns the contents of args in following order:
*	prettyprint flag
*	script flag
*	script path
*	command
*	planets
 */
func (opts *Opts) procArgs(args []string) {
	flag.BoolVar(&opts.helpFlag, "h", false, "help")
	flag.BoolVar(&opts.prettyFlag, "pp", false, "prettyprint")
	flag.BoolVar(&opts.typeFlag, "t", false, "type")
	flag.BoolVar(&opts.debugFlag, "d", false, "verbose")
	flag.BoolVar(&opts.loadFlag, "l", false, "ssh profile loading")
	flag.BoolVar(&opts.versionFlag, "v", false, "version")
	flag.StringVar(&opts.template, "tp", "", "filename of template")
	flag.StringVar(&opts.scriptName, "s", "", "name of the script(regardless wether db or bash) to be executed")
	flag.StringVar(&opts.command, "c", "", "command to be executed in quotes")
	flag.Parse()
	opts.scriptFlag = !(opts.scriptName == "")
	opts.tableFlag = !(opts.template == "")
	if opts.scriptFlag && !(opts.command == "") {
		var err error
		throwErrExt(err, "providing both a script AND a command is not possible")
	}

	planets := flag.Args()
	opts.command = strings.TrimSuffix(strings.TrimPrefix(opts.command, "\""), "\"")
	opts.template = strings.TrimSuffix(strings.TrimPrefix(opts.template, "\""), "\"")
	opts.scriptName = strings.TrimSuffix(strings.TrimPrefix(opts.scriptName, "\""), "\"")

	for _, argument := range planets {
		if isSupported(argument) {
			opts.planets = append(opts.planets, argument)
			opts.planetsCount++
		} else {
			fmt.Fprintf(os.Stderr, "This Type of Connection is not supported.")
			os.Exit(1)
		}
	}
	if len(args) == 1 {
		opts.helpFlag = true
	}

}

/**
*	Splits the given connectiondetails and returns the hostname
*	@params:
*		connDet: Connection details in following form: user@hostname
*	@return: hostname
 */
func getHost(connDet string) string {
	toReturn := strings.Split(connDet, "@")
	return toReturn[1]
}

/**
*	Splits the given connectiondetails and returns the user
*	@params:
*		connDet: Connection details in following form: user@hostname
*	@return: user
 */
func getUser(connDet string) string {
	toReturn := strings.Split(connDet, "@")
	return toReturn[0]
}

/**
*	Returns the type of a given planet
*	@params:
*		id: The planets id
*	@return: The planets type
 */
func getType(id string) string {
	cmd := exec.Command("ff", "-t", id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		throwErrOut(out, err)
	}
	return strings.TrimSpace(string(out))
}

/**
*	Returns the connection details to a given planet
*	@params:
*		id: The planets id
*	@return: The connection details to the planet
 */
func getPlanetDetails(id string) string {
	cmd := exec.Command("ff", id, "-f=pqdb")
	out, err := cmd.CombinedOutput()
	if err != nil {
		throwErrOut(out, err)
	}
	return strings.TrimSpace(string(out))
}

/**
*	counts the supported planets in a list of planets
 */
func countSupported(planets []string) int {
	i := 0
	for _, planet := range planets {
		if getType(planet) == linuxServer {
			i++
		}
	}
	return i
}

/**
*	checks, wether a planet is supported by goo or not
 */
func isSupported(planet string) bool {
	switch getType(planet) {
	case linuxServer:
		return true
	case database:
		return true
	case webServer:
		return false
	default:
		return false

	}

}

func getMaxLength(toProcess string) int {
	_ = toProcess
	return 0
}

/**
*	splits db details (dbID:user@host) and returns them as dbID,user@host
 */
func procDBDets(dbDet string) (string, string) {
	parts := strings.Split(dbDet, ":")
	return parts[0], parts[1]
}
