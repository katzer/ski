package main

import (
	"errors"
	"flag"
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

func (opts *Opts) procArgs(args []string) {
	flag.BoolVar(&opts.help, "h", false, "help")
	flag.BoolVar(&opts.pretty, "p", false, "prettyprint")
	flag.BoolVar(&opts.debug, "d", false, "verbose")
	flag.BoolVar(&opts.load, "l", false, "ssh profile loading")
	flag.BoolVar(&opts.version, "v", false, "version")
	flag.StringVar(&opts.template, "t", "", "filename of template")
	flag.StringVar(&opts.scriptName, "s", "", "name of the script(regardless wether db or bash) to be executed")
	flag.StringVar(&opts.command, "c", "", "command to be executed in quotes")
	flag.Parse()

	validateArgsCount(opts)
	if opts.version {
		printVersion()
		os.Exit(0)
	}

	validateCommandAndScript(opts.scriptName, opts.command)

	planets := flag.Args()
	opts.command = strings.TrimSuffix(strings.TrimPrefix(opts.command, "\""), "\"")
	opts.template = strings.TrimSuffix(strings.TrimPrefix(opts.template, "\""), "\"")
	opts.scriptName = strings.TrimSuffix(strings.TrimPrefix(opts.scriptName, "\""), "\"")

	for _, argument := range planets {
		opts.planets = append(opts.planets, argument)
	}
}

/**
*	Checks if theres a command and a script at the same time
 */
func validateCommandAndScript(scriptname string, command string) {
	if !(scriptname == "") && !(command == "") {
		message := "providing both a script AND a command is not possible"
		err := errors.New(message)
		os.Stderr.WriteString(fmt.Sprintf("%s\nAddInf: %s\n", err, message))
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
	versionNotWanted := !opts.version
	if tooFew || scriptEmpty && cmdEmpty && versionNotWanted {
		printUsage()
		os.Exit(0)
	}
}
