package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/**
################################################################################
						Information-Retrieval-Section
################################################################################
*/

type Opts struct {
	prettyFlag   bool
	scriptFlag   bool
	typeFlag     bool
	debugFlag    bool
	loadFlag     bool
	helpFlag     bool
	versionFlag  bool
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

	cleanArgs := args[1:]

	for _, argument := range cleanArgs {
		if strings.HasPrefix(argument, "-h") || strings.HasPrefix(argument, "--help") {
			opts.helpFlag = true
		} else if strings.HasPrefix(argument, "-p") || strings.HasPrefix(argument, "--pretty") {
			opts.prettyFlag = true
		} else if strings.HasPrefix(argument, "-t") || strings.HasPrefix(argument, "--type") {
			opts.typeFlag = true
		} else if strings.HasPrefix(argument, "-d") || strings.HasPrefix(argument, "--debug") {
			opts.debugFlag = true
		} else if strings.HasPrefix(argument, "-l") || strings.HasPrefix(argument, "--load") {
			opts.loadFlag = true
		} else if strings.HasPrefix(argument, "-v") || strings.HasPrefix(argument, "--version") {
			opts.versionFlag = true
		} else if strings.HasPrefix(argument, "-c") || strings.HasPrefix(argument, "--command") {
			// TODO what if theres a = in the command itself?
			opts.command = strings.TrimSuffix(strings.TrimPrefix(strings.Split(argument, "=")[1], "\""), "\"")
		} else if strings.HasPrefix(argument, "-s") || strings.HasPrefix(argument, "--script") {
			opts.scriptFlag = true
			opts.scriptPath = strings.Split(argument, "=")[1]
		} else {
			if isSupported(argument) {
				opts.planets = append(opts.planets, argument)
			} else {
				fmt.Fprintf(os.Stderr, "This Type of Connection is not supported.")
				os.Exit(1)
			}
		}
	}
	if len(args) < 3 {
		printHelp()
		os.Exit(0)
	}

	//return prettyFlag, scriptFlag, scriptPath, command, planets, debugFlag, typeFlag, loadFlag
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
func (opts Opts) getConnDet(id string) string {
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
	if getType(planet) == "server" {
		return true
	}
	return false

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
