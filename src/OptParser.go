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
	prettyFlag   bool
	scriptFlag   bool
	typeFlag     bool
	debugFlag    bool
	loadFlag     bool
	helpFlag     bool
	versionFlag  bool
	tableFlag    bool
	scriptPath   string
	command      string
	planets      []string
	planetsCount int
	currentDet   string
	currentDBDet string
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
	flag.BoolVar(&opts.tableFlag, "tp", false, "tablePrint")
	flag.StringVar(&opts.command, "c", "", "command to be executed in quotes")
	flag.StringVar(&opts.scriptPath, "s", "", "path to script file to be uploaded and executed")
	flag.Parse()
	opts.scriptFlag = !(opts.scriptPath == "")

	planets := flag.Args()
	opts.command = strings.TrimSuffix(strings.TrimPrefix(opts.command, "\""), "\"")
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
func getConnDet(id string) string {
	cmd := exec.Command("ff", id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		throwErrOut(out, err)
	}
	return strings.TrimSpace(string(out))
}

/**
*
*
*
 */
func countSupported(planets []string) int {
	i := 0
	for _, planet := range planets {
		if getType(planet) == "server" {
			i++
		}
	}
	return i
}

/**
*
*
 */
func isSupported(planet string) bool {
	switch getType(planet) {
	case "server":
		return true
	case "db":
		return true
	case "web":
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
*
 */
func procDBDets(dbDet string) (string, string) {
	parts := strings.Split(dbDet, ":")
	return parts[0], parts[1]
}
