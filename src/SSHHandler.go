package main

import (
	//"github.com/mgutz/ansi"
	"fmt"
	"os"
	"strings"

	"gopkg.in/hypersleep/easyssh.v0"
)

/**
*	Executes a command on a remote ssh server
*	@params:
*		connDet: connection details in following form: user@hostname
*		cmd: command to be executed
 */
func execCommand(user string, hostname string, command string, strucOut *StructuredOuput, opts *Opts) {

	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:    os.Getenv("ORBIT_KEY"),
		Port:   "22",
	}
	var cmd string
	if opts.loadFlag {
		cmd = fmt.Sprintf(`sh -lc "echo -----APPPLANT-ORBIT----- && %s "`, command)
	} else {
		cmd = command
	}
	// Call Run method with command you want to run on remote server.
	out, err := ssh.Run(cmd)
	// Handle errors
	if err != nil {
		throwErrExt(err, "called from execCommand ")
	} else {
		cleanedOut := out
		if opts.loadFlag {
			splitOut := strings.Split(out, "-----APPPLANT-ORBIT-----\n")
			cleanedOut = splitOut[len(splitOut)-1]
		}
		maxLength := getMaxLength(out)
		strucOut.output = cleanedOut
		strucOut.maxOutLength = maxLength
	}
	if opts.debugFlag {
		debugPrintString("### execCommand complete ###")
		debugPrintString(fmt.Sprintf("user: %s\n", user))
		debugPrintString(fmt.Sprintf("hostname: %s\n", hostname))
		debugPrintString(fmt.Sprintf("orbit key: %s\n", os.Getenv("ORBIT_KEY")))
		debugPrintString(fmt.Sprintf("command: %s\n", command))
		debugPrintString(fmt.Sprintf("strucOut: %v\n", strucOut))
		//debugPrintOpts(opts)
		debugPrintStructuredOutput(strucOut)
		debugPrintString("### execCommand complete ###")
	}
}

/**
*	Uploads a file to the remote server
 */
func uploadFile(user string, hostname string, opts *Opts) {
	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:    os.Getenv("ORBIT_KEY"),
		Port:   "22",
	}

	// Call Scp method with file you want to upload to remote server.
	err := ssh.Scp(opts.scriptPath)

	// Handle errors
	if err != nil {
		throwErrExt(err, "called from uploadFile")
	}
}

/**
*	Uploads and executes a script on a given planet
*	@params:
*		connDet: 	Connection details to planet
*		scriptPath: Path to script
 */
func execScript(user string, hostname string, strucOut *StructuredOuput, opts *Opts) {
	uploadFile(user, hostname, opts)
	path := strings.Split(opts.scriptPath, "/")
	placeholder := StructuredOuput{}
	scriptName := path[len(path)-1]
	executionCommand := fmt.Sprintf("chmod u+x %s && ./%s", scriptName, scriptName)
	delCommand := fmt.Sprintf("rm %s", scriptName)
	execCommand(user, hostname, executionCommand, strucOut, opts)
	execCommand(user, hostname, delCommand, &placeholder, opts)
}
