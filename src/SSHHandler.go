package main

import (
	//"github.com/mgutz/ansi"
	"gopkg.in/hypersleep/easyssh.v0"
	"os"
	"strings"
)

/**
################################################################################
								SSH-Section
################################################################################
*/

type SSHHandler struct {
}

/**
*	Executes a command on a remote ssh server
*	@params:
*		connDet: connection details in following form: user@hostname
*		cmd: command to be executed
 */
func execSSHCommand(connDet string, command string, strucOut *StructuredOuput, opts *Opts) {

	user := getUser(opts.currentDet)
	hostname := getHost(opts.currentDet)
	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:    os.Getenv("ORBIT_KEY"),
		Port:   "22",
	}
	var cmd string
	if opts.loadFlag {
		cmd = "sh -lc \"echo -----APPPLANT-ORBIT----- && " + command + " \""
	} else {
		cmd = command
	}
	// Call Run method with command you want to run on remote server.
	out, err := ssh.Run(cmd)
	// Handle errors
	if err != nil {
		throwErr(err)
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
}

/**
*	Uploads a file to the remote server
 */
func uploadFile(connDet string, opts *Opts) {
	user := getUser(opts.currentDet)
	hostname := getHost(opts.currentDet)

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
		throwErr(err)
	}
}

/**
*	Uploads and executes a script on a given planet
*	@params:
*		connDet: 	Connection details to planet
*		scriptPath: Path to script
 */
func upAndExecSSHScript(connDet string, strucOut *StructuredOuput, opts *Opts) {
	uploadFile(connDet, opts)
	path := strings.Split(opts.scriptPath, "/")
	placeholder := StructuredOuput{}
	scriptName := path[len(path)-1]
	execSSHCommand(connDet, "chmod +x "+scriptName+" && "+"./"+scriptName, &placeholder, opts)
}
