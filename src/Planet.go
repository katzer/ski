package main

import "fmt"
import "strings"

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
			if !strings.HasSuffix(opts.scriptName, ".sql") {
				opts.scriptName = fmt.Sprintf("%s.sql", opts.scriptName)
			}
			execDBScript(planet.dbID, planet.user, planet.host, &planet.outputStruct, opts)
		} else {
			execDBCommand(planet.dbID, planet.user, planet.host, &planet.outputStruct, opts)
		}
		//trimDBMetaInformations(&planet.outputStruct)
	} else if planet.planetType == linuxServer {
		if opts.scriptFlag {
			if !strings.HasSuffix(opts.scriptName, ".sh") {
				opts.scriptName = fmt.Sprintf("%s.sh", opts.scriptName)
			}
			execScript(planet.user, planet.host, &planet.outputStruct, opts)
		} else {
			if opts.debugFlag {
				fmt.Println("###planet.execute-->execcommand###")
				fmt.Printf("planet.user: %s\n", planet.user)
				fmt.Printf("planet.host: %s\n", planet.host)
				fmt.Printf("opts.command: %s\n", opts.command)
				fmt.Printf("planet.outputStruct: %v\n", planet.outputStruct)
				fmt.Printf("opts: %v\n", opts)
				fmt.Println("###planet.execute-->execcommand###")
			}
			execCommand(planet.user, planet.host, opts.command, &(planet.outputStruct), opts)
			if opts.debugFlag {
				fmt.Println("###planet.execute execcommand-->###")
				fmt.Printf("planet.user: %s\n", planet.user)
				fmt.Printf("planet.host: %s\n", planet.host)
				fmt.Printf("opts.command: %s\n", opts.command)
				fmt.Printf("planet.outputStruct: %v\n", planet.outputStruct)
				fmt.Printf("opts: %v\n", opts)
				fmt.Println("###planet.execute execcommand-->###")
			}
		}
	} else {

	}
}
