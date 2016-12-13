package main

import (
	"os"
	"os/exec"
	"strings"
)

/**
################################################################################
						Information-Retrieval-Section
################################################################################
*/

/**
*	Returns the contents of args in following order:
*	prettyprint flag
*	script flag
*	script path
*	command
*	planets
 */
func procArgs(args []string) (bool, bool, string, string, []string, bool, bool) {
	prettyFlag := false
	scriptFlag := false
	typeFlag := false
	debugFlag := false
	var scriptPath string = ""
	var command string = ""
	planets := make([]string, 0)

	cleanArgs := args[1:]

	for _, argument := range cleanArgs {
		if strings.HasPrefix(argument, "-h") || strings.HasPrefix(argument, "--help") {
			printHelp()
			os.Exit(0)
		} else if strings.HasPrefix(argument, "-p") || strings.HasPrefix(argument, "--pretty") {
			prettyFlag = true
		} else if strings.HasPrefix(argument, "-t") || strings.HasPrefix(argument, "--type") {
			typeFlag = true
		} else if strings.HasPrefix(argument, "-d") || strings.HasPrefix(argument, "--debug") {
			debugFlag = true
		} else if strings.HasPrefix(argument, "-v") || strings.HasPrefix(argument, "--version") {
			printVersion()
			os.Exit(0)
		} else if strings.HasPrefix(argument, "-c") || strings.HasPrefix(argument, "--command") {
			// TODO what if theres a = in the command itself?
			command = strings.TrimSuffix(strings.TrimPrefix(strings.Split(argument, "=")[1], "\""), "\"")
		} else if strings.HasPrefix(argument, "-s") || strings.HasPrefix(argument, "--script") {
			scriptFlag = true
			scriptPath = strings.Split(argument, "=")[1]
		} else {
			planets = append(planets, argument)
		}
	}
	if len(args) < 3 {
		printHelp()
		os.Exit(0)
	}

	_ = prettyFlag

	return prettyFlag, scriptFlag, scriptPath, command, planets, debugFlag, typeFlag
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
*					DEPRECATED
*
*	Extracts the desired argument from the arguments list.
*	@params:
*		args: Arguments to be searched in.
*		type: Type of desired Argument (command,id)
*		position: starting position of desired argument
*	@return: The desired arguments
 */
func getArg(args []string, argType string, position int) string {
	switch argType {
	case "command":
		var command string = args[position]
		var cmdArgs []string
		if len(args) > (position + 1) {
			cmdArgs = args[(position + 1):(len(args))]
			for _, argument := range cmdArgs {
				command += (" " + argument)
			}
		}
		return command
	default:
		return args[position]
	}

}
