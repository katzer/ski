package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"strings"

	log "github.com/Sirupsen/logrus"
)

func execDBCommand(planet *Planet, opts *Opts) error {
	var err error
	if !strings.HasSuffix(opts.Command, ";") {
		log.Warningf("The SQL-Command needs to be terminated with a \";\"")
		log.Warningf("Appending \";\"...")
		opts.Command = fmt.Sprintf("%s;", opts.Command)
	}
	tmpDBFile := path.Join(os.Getenv("ORBIT_HOME"), "scripts", "orbit.sql")
	err = ioutil.WriteFile(tmpDBFile, []byte(opts.Command), 0644)
	if err != nil {
		errormessage := fmt.Sprintf("writing temporary sql script failed : %v", err)
		planet.outputStruct.output = fmt.Sprintf("%s\n%s", planet.outputStruct.output, errormessage)
		planet.errored = true
		log.Warning(errormessage)
		return err
	}
	opts.ScriptName = "orbit.sql"
	err = execDBScript(planet, opts)
	if err != nil {
		return err
	}
	err = os.Remove(tmpDBFile)
	if err != nil {
		errormessage := fmt.Sprintf("removing temporary sql script failed : %v", err)
		planet.outputStruct.output = fmt.Sprintf("%s\n%s", planet.outputStruct.output, errormessage)
		planet.errored = true
		log.Warning(errormessage)
		return err
	}
	return err
}

func execDBScript(planet *Planet, opts *Opts) error {
	var err error
	err = uploadFile(planet, opts)
	if err != nil {
		return err
	}
	scriptName := opts.ScriptName
	moveCommand := fmt.Sprintf("mv ~/%s ~/sql/%s>/dev/null", scriptName, scriptName)
	err = execCommand(moveCommand, planet, opts)
	if err != nil {
		return err
	}
	queryString := fmt.Sprintf(dbCommand, planet.user, planet.dbID, scriptName)
	err = execCommand(queryString, planet, opts)
	if err != nil {
		return err
	}
	removeCommand := fmt.Sprintf("rm ~/sql/%s>/dev/null", scriptName)
	err = execCommand(removeCommand, planet, opts)
	if err != nil {
		return err
	}
	cleanDBMetaData(planet.outputStruct)
	return err
}

func cleanDBMetaData(strucOut *StructuredOuput) {
	split := strings.Split(strucOut.output, "\n")
	reduced := split[1:(len(split) - 3)]
	strucOut.output = strings.Join(reduced, "\n")
}
