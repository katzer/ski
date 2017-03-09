package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"strings"
)

func execDBCommand(planet *Planet, strucOut *StructuredOuput, opts *Opts) {
	if !strings.HasSuffix(opts.Command, ";") {
		log.Warningf("The SQL-Command needs to be terminated with a \";\"")
		log.Warningf("Appending \";\"...")
		opts.Command = fmt.Sprintf("%s;", opts.Command)
	}
	tmpDBFile := path.Join(os.Getenv("ORBIT_HOME"), "scripts", "orbit.sql")
	err := ioutil.WriteFile(tmpDBFile, []byte(opts.Command), 0644)
	if err != nil {
		log.Fatalf("writing temporary sql script failed : %v", err)
	}
	opts.ScriptName = "orbit.sql"
	execDBScript(planet, strucOut, opts)
	err = os.Remove(tmpDBFile)
	if err != nil {
		log.Fatal(err)
	}
}

func execDBScript(planet *Planet, strucOut *StructuredOuput, opts *Opts) {
	const dbCommand = ". profiles/%s.prof && pqdb_sql.out -x -s %s ~/sql/%s"
	uploadFile(planet.user, planet.host, opts)
	placeholder := StructuredOuput{}
	scriptName := opts.ScriptName
	command := fmt.Sprintf("mv ~/%s ~/sql/%s", scriptName, scriptName)
	execCommand(command, planet, &placeholder, opts)
	queryString := fmt.Sprintf(dbCommand, planet.user, planet.dbID, scriptName)
	removeCommand := fmt.Sprintf("rm ~/sql/%s", scriptName)
	execCommand(queryString, planet, strucOut, opts)
	execCommand(removeCommand, planet, &placeholder, opts)
	cleanDBMetaData(strucOut)
}

func cleanDBMetaData(strucOut *StructuredOuput) {
	split := strings.Split(strucOut.output, "\n")
	reduced := split[1:(len(split) - 3)]
	strucOut.output = strings.Join(reduced, "\n")
}
