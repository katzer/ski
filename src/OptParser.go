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

	validateCommandAndScript(opts.scriptName, opts.command)

	planets := flag.Args()
	opts.command = strings.TrimSuffix(strings.TrimPrefix(opts.command, "\""), "\"")
	opts.template = strings.TrimSuffix(strings.TrimPrefix(opts.template, "\""), "\"")
	opts.scriptName = strings.TrimSuffix(strings.TrimPrefix(opts.scriptName, "\""), "\"")

	for _, argument := range planets {
		opts.planets = append(opts.planets, argument)
	}

	validateArgsCount(opts)

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
	if len(os.Args) == 1 {
		opts.help = true
	}
	if !opts.help && opts.scriptName == "" && !opts.version && opts.command == "" {
		opts.help = true
	}
}
