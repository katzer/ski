package main

import (
	"errors"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"

	"strings"
)

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

	planets := flag.Args()
	opts.command = strings.TrimSuffix(strings.TrimPrefix(opts.command, "\""), "\"")
	opts.template = strings.TrimSuffix(strings.TrimPrefix(opts.template, "\""), "\"")
	opts.scriptName = strings.TrimSuffix(strings.TrimPrefix(opts.scriptName, "\""), "\"")

	for _, argument := range planets {
		opts.planets = append(opts.planets, argument)
		// opts.planetsCount++
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
