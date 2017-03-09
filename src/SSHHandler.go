package main

import (
	"fmt"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/hypersleep/easyssh.v0"
)

func execCommand(command string, planet *Planet, opts *Opts) error {
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
		errorString := fmt.Sprintf("%s\nAdditional Info: %s\n", err, message)
		log.Warn(errorString)
		planet.outputStruct.output = fmt.Sprintf("%s\n%s", planet.outputStruct.output, errorString)
		planet.errored = true
		logExecCommand(command, planet)
		return err
	}
	out = cleanProfileLoadedOutput(out, opts)

	planet.outputStruct.output += out
	logExecCommand(command, planet)
	return nil
}

func uploadFile(planet *Planet, opts *Opts) error {
	keyPath := getKeyPath()

	ssh := &easyssh.MakeConfig{
		User:   planet.user,
		Server: planet.host,
		Key:    keyPath,
		Port:   "22",
	}

	// Call Scp method with file you want to upload to remote server.
	err := ssh.Scp(path.Join(os.Getenv("ORBIT_HOME"), scriptDirectory, opts.ScriptName))

	// Handle errors
	if err != nil {
		message := fmt.Sprintf("called from uploadFile. Keypath: %s", keyPath)
		errorString := fmt.Sprintf("%s\nAddInf: %s\n", err, message)
		log.Warn(errorString)
		planet.outputStruct.output = fmt.Sprintf("%s\n%s", planet.outputStruct.output, errorString)
		planet.errored = true
		return err
	}
	return nil
}

func execScript(planet *Planet, opts *Opts) error {
	var err error
	log.Debugf("function: execScript")
	log.Debugf("user, host : |%s| |%s|", planet.user, planet.host)
	err = uploadFile(planet, opts)
	if err != nil {
		return err
	}
	scriptName := opts.ScriptName
	executionCommand := fmt.Sprintf("sh %s", scriptName)
	delCommand := fmt.Sprintf("rm %s>/dev/null", scriptName)
	err = execCommand(executionCommand, planet, opts)
	if err != nil {
		return err
	}
	err = execCommand(delCommand, planet, opts)
	if err != nil {
		return err
	}
	return nil
}
