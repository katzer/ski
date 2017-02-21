package main

import (
	log "github.com/Sirupsen/logrus"
)

// Planet contains all Informations of one server
type Planet struct {
	id           string
	user         string
	host         string
	planetType   string
	dbID         string
	outputStruct StructuredOuput
}

func (planet *Planet) execute(opts *Opts) {
	if planet.planetType == database {
		planet.executeDatabase(opts)
	} else if planet.planetType == linuxServer {
		planet.executeLinux(opts)
	}
}

func (planet *Planet) executeDatabase(opts *Opts) {
	if opts.scriptName != "" {
		execDBScript(planet, &planet.outputStruct, opts)
	} else {
		execDBCommand(planet, &planet.outputStruct, opts)
	}
}

func (planet *Planet) executeLinux(opts *Opts) {
	if opts.scriptName != "" {
		execScript(planet, &planet.outputStruct, opts)
	} else {
		planet.planetInfo(opts)
		execCommand(opts.command, planet, &(planet.outputStruct), opts)
		planet.planetInfo(opts)
	}
}

func (planet *Planet) planetInfo(opts *Opts) {
	log.Debugln("###planet.execute-->execcommand###")
	log.Debugln("planet.user: %s", planet.user)
	log.Debugln("planet.host: %s", planet.host)
	log.Debugln("opts.command: %s", opts.command)
	log.Debugln("planet.outputStruct: %v", planet.outputStruct)
	log.Debugln("opts: %v\n", opts)
	log.Debugln("###planet.execute-->execcommand###")
}
