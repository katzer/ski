package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// codebeat:disable[TOO_MANY_IVARS]

// Opts structure for holding commandline arguments
type Opts struct {
	Debug      bool `json:"debug"`
	Help       bool `json:"help"`
	Load       bool `json:"load"`
	Pretty     bool
	Version    bool     `json:"version"`
	MaxToKeep  int      `json:"max_to_keep"`
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
	MaxToKeep: %d
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
		opts.MaxToKeep,
		opts.Command,
		opts.ScriptName,
		opts.Template,
		opts.Planets)
}

func (opts *Opts) validate() {
	opts.validateExtension()
	opts.validateCommandAndScript()
	opts.checkForInvalidIds()
	opts.validateTemplate()
}

func (opts *Opts) checkForInvalidIds() {
	for _, id := range opts.Planets {
		// Check if any flags were given after planet ids, if yes stop the app
		if strings.HasPrefix(id, "-") {
			fmt.Fprintf(os.Stderr, "Unknown target: %s", id)
			os.Exit(1)
		}
	}
}

func (opts *Opts) validateCommandAndScript() {
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
func (opts *Opts) validateExtension() {
	script := opts.ScriptName
	isSH := strings.HasSuffix(script, ".sh")
	isSQL := strings.HasSuffix(script, ".sql")
	if script != "" && !(isSH || isSQL) {
		message := fmt.Sprintf("The provided scripts %s must have either .sh or .sql extension.", script)
		fmt.Fprintln(os.Stderr, message)
		log.Fatal(message)
	}
}

func (opts *Opts) validateTemplate() {
	template := opts.Template != ""
	file := path.Join(os.Getenv("ORBIT_HOME"), "templates", opts.Template)
	if template {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			message := fmt.Sprintf("The provided template does not exist")
			fmt.Fprintln(os.Stderr, message)
			log.Fatal(message)
		}
	}
}
