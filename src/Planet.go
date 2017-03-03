package main

import (
	log "github.com/Sirupsen/logrus"
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
	outputStruct StructuredOuput
}

// StructuredOuput ...
type StructuredOuput struct {
	planet       string
	output       string
	maxOutLength int
}
// codebeat:enable[TOO_MANY_IVARS]

func (planet *Planet) execute(opts *Opts) {
	if planet.planetType == database {
		planet.executeDatabase(opts)
	} else if planet.planetType == linuxServer {
		planet.executeLinux(opts)
	}
}

func (planet *Planet) executeDatabase(opts *Opts) {
	if opts.ScriptName != "" {
		execDBScript(planet, &planet.outputStruct, opts)
	} else {
		execDBCommand(planet, &planet.outputStruct, opts)
	}
}

func (planet *Planet) executeLinux(opts *Opts) {
	if opts.ScriptName != "" {
		execScript(planet, &planet.outputStruct, opts)
	} else {
		planet.planetInfo(opts)
		execCommand(opts.Command, planet, &(planet.outputStruct), opts)
		planet.planetInfo(opts)
	}
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
