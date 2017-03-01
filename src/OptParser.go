package main

import (
	"errors"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"

	"strings"
)

// codebeat:disable[TOO_MANY_IVARS]
// Opts ...
type Opts struct {
	debugFlag   bool
	helpFlag    bool
	loadFlag    bool
	prettyFlag  bool
	versionFlag bool
	command     string
	scriptName  string
	template    string
	planets     []string
}

// codebeat:enable[TOO_MANY_IVARS]

func (opts *Opts) String() string {
	template := `opts : {
	debugFlag: %t
	helpFlag: %t
	loadFlag: %t
	prettyFlag: %t
	versionFlag: %t
	command: %s
	scriptName: %s
	planets: %v
}
`

	return fmt.Sprintf(template,
		opts.debugFlag,
		opts.helpFlag,
		opts.loadFlag,
		opts.prettyFlag,
		opts.versionFlag,
		opts.command,
		opts.scriptName,
		opts.planets)
}

func (opts *Opts) procArgs(args []string) {
	flag.BoolVar(&opts.helpFlag, "h", false, "help")
	flag.BoolVar(&opts.prettyFlag, "p", false, "prettyprint")
	flag.BoolVar(&opts.debugFlag, "d", false, "verbose")
	flag.BoolVar(&opts.loadFlag, "l", false, "ssh profile loading")
	flag.BoolVar(&opts.versionFlag, "v", false, "version")
	flag.StringVar(&opts.template, "t", "", "filename of template")
	flag.StringVar(&opts.scriptName, "s", "", "name of the script(regardless wether db or bash) to be executed")
	flag.StringVar(&opts.command, "c", "", "command to be executed in quotes")
	flag.Parse()

	validateCommandAndScript(opts.scriptName, opts.command)
	validateExtension(opts.scriptName)

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
		opts.helpFlag = true
	}
	if !opts.helpFlag && opts.scriptName == "" && !opts.versionFlag && opts.command == "" {
		opts.helpFlag = true
	}
}

func validateExtension(scriptname string) {
	if scriptname != "" && !(strings.HasSuffix(scriptname, ".sh") || strings.HasSuffix(scriptname, ".sql")) {
		message := "The provided script must end in either \".sh\" if it's a shell script or \".sql\" if it's a sql script ."
		err := errors.New(message)
		os.Stderr.WriteString(message)
		log.Fatal(err)
	}
}
