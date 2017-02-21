package main

import (
	"fmt"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/hypersleep/easyssh.v0"
)

/**
*	Executes a command on a remote ssh server
*	@params:
*		connDet: connection details in following form: user@hostname
*		cmd: command to be executed
 */
func execCommand(command string, planet *Planet, strucOut *StructuredOuput, opts *Opts) {
	keyPath := getKeyPath()

	ssh := &easyssh.MakeConfig{
		User:   planet.user,
		Server: planet.host,
		Key:    keyPath,
		Port:   "22",
	}
	cmd := makeLoadCommand(command, opts)
	// Call Run method with command you want to run on remote server.
	out, err := ssh.Run(cmd)
	// Handle errors
	if err != nil {
		message := fmt.Sprintf("called from execCommand.\nKeypath: %s\nCommand: %s", keyPath, cmd)
		errorString := fmt.Sprintf("%s\nAddInf: %s\n", err, message)
		os.Stderr.WriteString(errorString)
		log.Fatalf("%s\nAdditional info: %s\n", err, message)
	}
	cleanedOut := cleanProfileLoadedOutput(out, opts)
	strucOut.output = cleanedOut
	strucOut.maxOutLength = 0
	logExecCommand(command, planet, strucOut)
}

/**
*	Uploads a file to the remote server
 */
func uploadFile(user string, hostname string, opts *Opts) {
	keyPath := getKeyPath()

	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:    keyPath,
		Port:   "22",
	}

	// Call Scp method with file you want to upload to remote server.
	err := ssh.Scp(path.Join(os.Getenv("ORBIT_HOME"), scriptDirectory, opts.scriptName))

	// Handle errors
	if err != nil {
		message := fmt.Sprintf("called from uploadFile. Keypath: %s", keyPath)
		errorString := fmt.Sprintf("%s\nAddInf: %s\n", err, message)
		os.Stderr.WriteString(errorString)
		log.Fatalf("%s\nAdditional info: %s\n", err, message)
	}
}

/**
*	Uploads and executes a script on a given planet
*	@params:
*		connDet: 	Connection details to planet
*		scriptPath: Path to script
 */
func execScript(planet *Planet, strucOut *StructuredOuput, opts *Opts) {
	uploadFile(planet.user, planet.host, opts)
	placeholder := StructuredOuput{}
	scriptName := opts.scriptName
	executionCommand := fmt.Sprintf("sh %s", scriptName)
	delCommand := fmt.Sprintf("rm %s", scriptName)
	execCommand(executionCommand, planet, strucOut, opts)
	execCommand(delCommand, planet, &placeholder, opts)
}
