package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"strings"

	log "github.com/Sirupsen/logrus"
)

func execDBCommand(planet *Planet, opts *Opts) {
	if !strings.HasSuffix(opts.Command, ";") {
		log.Warningf("The SQL-Command needs to be terminated with a \";\"")
		log.Warningf("Appending \";\"...")
		opts.Command = fmt.Sprintf("%s;", opts.Command)
	}
	tmpDBFile := path.Join(os.Getenv("ORBIT_HOME"), "scripts", "orbit.sql")
	err := ioutil.WriteFile(tmpDBFile, []byte(opts.Command), 0644)
	if err != nil {
		errormessage := fmt.Sprintf("writing temporary sql script failed : %v", err)
		planet.outputStruct.output = fmt.Sprintf("%s\n%s", planet.outputStruct.output, errormessage)
		planet.errored = true
		log.Warning(errormessage)
	}
	opts.ScriptName = "orbit.sql"
	execDBScript(planet, opts)
	err = os.Remove(tmpDBFile)
	if err != nil {
		errormessage := fmt.Sprintf("removing temporary sql script failed : %v", err)
		planet.outputStruct.output = fmt.Sprintf("%s\n%s", planet.outputStruct.output, errormessage)
		planet.errored = true
		log.Warning(errormessage)
	}
}

func execDBScript(planet *Planet, opts *Opts) {
	uploadFile(planet, opts)
	scriptName := opts.ScriptName
	moveCommand := fmt.Sprintf("mv ~/%s ~/sql/%s>/dev/null", scriptName, scriptName)
	execCommand(moveCommand, planet, opts)
	queryString := fmt.Sprintf(dbCommand, planet.user, planet.dbID, scriptName)
	removeCommand := fmt.Sprintf("rm ~/sql/%s>/dev/null", scriptName)
	execCommand(queryString, planet, opts)
	execCommand(removeCommand, planet, opts)
	cleanDBMetaData(planet.outputStruct)
}

func cleanDBMetaData(strucOut *StructuredOuput) {
	split := strings.Split(strucOut.output, "\n")
	reduced := split[1:(len(split) - 3)]
	strucOut.output = strings.Join(reduced, "\n")
}
