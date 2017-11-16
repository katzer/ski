package main

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

// codebeat:disable[TOO_MANY_IVARS]

// Planet contains all Informations of one server
type Planet struct {
	id           string
	name         string
	user         string
	host         string
	planetType   string
	dbID         string
	valid        bool
	outputStruct *StructuredOuput
}

// StructuredOuput ...
type StructuredOuput struct {
	planet   string
	output   string
	keys     []string
	table    [][]string
	position int
	errored  bool
	errors   map[string]string
}

// codebeat:enable[TOO_MANY_IVARS]

func (planet *Planet) execute(opts *Opts) {

	if !planet.validateType(opts) {
		planet.outputStruct.output = "script type doesn't fit to planetType"
		return
	}

	if planet.planetType == database {
		planet.executeDatabase(opts)
	} else if planet.planetType == linuxServer {
		planet.executeLinux(opts)
	}
}

func (planet *Planet) executeDatabase(opts *Opts) {
	if opts.ScriptName != "" {
		execDBScript(planet, opts)
	} else {
		execDBCommand(planet, opts)
	}
}

func (planet *Planet) executeLinux(opts *Opts) {
	if opts.ScriptName != "" {
		execScript(planet, opts)
	} else {
		execCommand(opts.Command, planet, opts)
	}
}

func (planet *Planet) validateType(opts *Opts) bool {
	if opts.ScriptName == "" {
		return true
	}
	sql := strings.HasSuffix(strings.ToLower(opts.ScriptName), ".sql")
	sh := strings.HasSuffix(strings.ToLower(opts.ScriptName), ".sh")
	if (sh) && planet.planetType == "server" {
		return true
	}
	if (sql) && planet.planetType == "db" {
		return true
	}
	return false
}

func (planet *Planet) planetInfo(opts *Opts) {
	log.Debugln("###planet.execute-->execcommand###")
	log.Debugln("planet.user: %s", planet.user)
	log.Debugln("planet.host: %s", planet.host)
	log.Debugln("opts.command: %s", opts.Command)
	log.Debugln("planet.outputStruct: %v", planet.outputStruct)
	log.Debugln("opts: %v\n", opts)
	log.Debugln("###planet.execute-->execcommand###")
}
