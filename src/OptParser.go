package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"

	"strings"
)

// Opts ...
type Opts struct {
	debug      bool
	help       bool
	load       bool
	pretty     bool
	version    bool
	command    string
	scriptName string
	template   string
	planets    []string
}

// Task ...
// Note: will replace Opts above or Opts will look like this.
// TODO: refactoring pending.
type Task struct {
	Debug      bool     `json:"debug"`
	Help       bool     `json:"help"`
	Load       bool     `json:"load"`
	Pretty     bool     `json:"pretty"`
	Version    bool     `json:"version"`
	Command    string   `json:"command"`
	ScriptName string   `json:"scriptName"`
	Template   string   `json:"template"`
	Planets    []string `json:"planets"`
}

func (opts *Opts) String() string {
	template := `opts : {
	debug: %t
	help: %t
	load: %t
	pretty: %t
	version: %t
	command: %s
	scriptName: %s
	planets: %v
}
`

	return fmt.Sprintf(template,
		opts.debug,
		opts.help,
		opts.load,
		opts.pretty,
		opts.version,
		opts.command,
		opts.scriptName,
		opts.planets)
}

func (opts *Opts) postProcessing() {

	opts.command = strings.TrimSuffix(strings.TrimPrefix(opts.command, "\""), "\"")
	opts.template = strings.TrimSuffix(strings.TrimPrefix(opts.template, "\""), "\"")
	opts.scriptName = strings.TrimSuffix(strings.TrimPrefix(opts.scriptName, "\""), "\"")

	// planets := flag.Args()
	// for _, argument := range planets {
	// 	opts.planets = append(opts.planets, argument)
	// }
}

func validate(opts *Opts) {
	validateArgsCount(opts)
	validateCommandAndScript(opts)
}

/**
*	Checks if theres a command and a script at the same time
 */
func validateCommandAndScript(opts *Opts) {
	scriptNotEmpty := len(opts.scriptName) > 0
	cmdNotEmpty := len(opts.command) > 0
	if scriptNotEmpty && cmdNotEmpty {
		message := "providing both a script AND a command is not possible"
		err := errors.New(message)
		fmt.Fprintf(os.Stderr, "%s\nAddInf: %s\n", err, message)
		log.Fatal(err)
	}
}

/**
*	Checks if there are enough of the correct arguments to run ski
 */
func validateArgsCount(opts *Opts) {
	tooFew := len(os.Args) == 1
	// TODO Check if flags package removes the leading and trailing white spaces.
	scriptEmpty := len(opts.scriptName) == 0
	cmdEmpty := len(opts.command) == 0
	if opts.version {
		printVersion()
		os.Exit(0)
	}
	if tooFew || scriptEmpty && cmdEmpty {
		printUsage()
		os.Exit(0)
	}
}
