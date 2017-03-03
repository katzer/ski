package main

import (
	"fmt"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/hypersleep/easyssh.v0"
)

func execCommand(command string, planet *Planet, strucOut *StructuredOuput, opts *Opts) {
	log.Debugf("function: execCommand")
	log.Debugf("user, host : %s %s", planet.user, planet.host)
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
		log.Warnf("%s\nAdditional info: %s\n", err, message)
		strucOut.output = message
		logExecCommand(command, planet, strucOut)
		return
	}
	cleanedOut := cleanProfileLoadedOutput(out, opts)
	strucOut.output = cleanedOut
	logExecCommand(command, planet, strucOut)
}

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
		log.Warningf("%s\nAdditional info: %s\n", err, message)
	}
}

func execScript(planet *Planet, strucOut *StructuredOuput, opts *Opts) {
	log.Debugf("function: execScript")
	log.Debugf("user, host : |%s| |%s|", planet.user, planet.host)
	uploadFile(planet.user, planet.host, opts)
	placeholder := StructuredOuput{}
	scriptName := opts.scriptName
	executionCommand := fmt.Sprintf("sh %s", scriptName)
	delCommand := fmt.Sprintf("rm %s", scriptName)
	execCommand(executionCommand, planet, strucOut, opts)
	execCommand(delCommand, planet, &placeholder, opts)
}
