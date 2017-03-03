package main

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

// codebeat:disable[TOO_MANY_IVARS]

// Opts structure for holding commandline arguments
type Opts struct {
	Debug      bool     `json:"debug"`
	Help       bool     `json:"help"`
	Load       bool     `json:"load"`
	Pretty     bool     `json:"pretty"`
	Version    bool     `json:"version"`
	SaveReport bool     `json:"save_report"`
	Command    string   `json:"command"`
	ScriptName string   `json:"scriptName"`
	Template   string   `json:"template"`
	Planets    []string `json:"planets"`
	LogFile    string   `json:"log_file"`
}

// codebeat:enable[TOO_MANY_IVARS]

func (opts *Opts) String() string {
	template := `opts : {
	Debug: %t
	Help: %t
	Load: %t
	Pretty: %t
	Version: %t
	Command: %s
	ScriptName: %s
	Template : %s
	Planets: %v
}
`

	return fmt.Sprintf(template,
		opts.Debug,
		opts.Help,
		opts.Load,
		opts.Pretty,
		opts.Version,
		opts.Command,
		opts.ScriptName,
		opts.Template,
		opts.Planets)
}

func (opts *Opts) postProcessing() {
	// TODO ask what is happening here
	opts.Command = strings.TrimSuffix(strings.TrimPrefix(opts.Command, "\""), "\"")
	opts.Template = strings.TrimSuffix(strings.TrimPrefix(opts.Template, "\""), "\"")
	opts.ScriptName = strings.TrimSuffix(strings.TrimPrefix(opts.ScriptName, "\""), "\"")
}

func validate(opts *Opts) {
	validateArgsCount(opts)
	validateExtension(opts.ScriptName)
	validateCommandAndScript(opts)
}

/**
*	Checks if theres a command and a script at the same time
 */
func validateCommandAndScript(opts *Opts) {
	scriptNotEmpty := len(opts.ScriptName) > 0
	cmdNotEmpty := len(opts.Command) > 0
	if scriptNotEmpty && cmdNotEmpty {
		message := "providing both a script AND a command is not possible"
		err := errors.New(message)
		fmt.Fprintf(os.Stderr, "%s\nAddInf: %s\n", err, message)
		log.Fatal(err)
	}
}

//Checks if the given script got an acceptable extension
func validateExtension(scriptname string) {
	if scriptname != "" && !(strings.HasSuffix(scriptname, ".sh") || strings.HasSuffix(scriptname, ".sql")) {
		message := fmt.Sprintf("The provided script \"%s\" must end in either \".sh\" if it's a shell script or \".sql\" if it's a sql script .", scriptname)
		err := errors.New(message)
		os.Stderr.WriteString(message)
		log.Fatal(err)
	}
}

/**
*	Checks if there are enough of the correct arguments to run ski
 */
func validateArgsCount(opts *Opts) {
	tooFew := len(os.Args) == 1
	// TODO Check if flags package removes the leading and trailing white spaces.
	scriptEmpty := len(opts.ScriptName) == 0
	cmdEmpty := len(opts.Command) == 0
	if opts.Version {
		printVersion()
		os.Exit(0)
	}
	if tooFew || scriptEmpty && cmdEmpty {
		printUsage()
		os.Exit(0)
	}
}

// creates a task from a json file
func createATaskFromJobFile(jsonFile string) (opts Opts) {
	job := Opts{}
	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Couldn't open job file: %s", jsonFile)
	}

	err = json.Unmarshal(bytes, &job)
	if err != nil {
		log.Fatalf("Error parsing job json file: %s", jsonFile)
	}

	log.Debugf("Read a task from %s", jsonFile)
	log.Debugln()
	log.Debugf("Unmarshalled %v", job)
	log.Debugln()
	return job
}
