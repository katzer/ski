package main

import "fmt"
import "strings"
import log "github.com/Sirupsen/logrus"

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
		if opts.scriptFlag {
			execDBScript(planet.dbID, planet.user, planet.host, &planet.outputStruct, opts)
		} else {
			execDBCommand(planet.dbID, planet.user, planet.host, &planet.outputStruct, opts)
		}
		//trimDBMetaInformations(&planet.outputStruct)
	} else if planet.planetType == linuxServer {
		if opts.scriptFlag {
			execScript(planet.user, planet.host, &planet.outputStruct, opts)
		} else {
			planet.planetInfo(opts)
			execCommand(planet.user, planet.host, opts.command, &(planet.outputStruct), opts)
			planet.planetInfo(opts)
		}
	} else {
		// TODO: Huh?
	}
}

func (planet *Planet) planetInfo(opts *Opts){
	log.Debugln("###planet.execute-->execcommand###")
	log.Debugln("planet.user: %s", planet.user)
	log.Debugln("planet.host: %s", planet.host)
	log.Debugln("opts.command: %s", opts.command)
	log.Debugln("planet.outputStruct: %v", planet.outputStruct)
	log.Debugln("opts: %v\n", opts)
	log.Debugln("###planet.execute-->execcommand###")
}
