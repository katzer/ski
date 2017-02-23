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
	if !strings.HasSuffix(opts.command, ";") {
		log.Fatal("The SQL-Command needs to be terminated with a \";\"")
	}
	tmpDBFile := path.Join(os.Getenv("ORBIT_HOME"), "scripts", "orbit.sql")
	err := ioutil.WriteFile(tmpDBFile, []byte(opts.command), 0644)
	if err != nil {
		log.Fatalf("writing temporary sql script failed : %v", err)
	}
	opts.scriptName = "orbit.sql"
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
	scriptName := opts.scriptName
	command := fmt.Sprintf("mv ~/%s ~/sql/%s", scriptName, scriptName)
	execCommand(command, planet, &placeholder, opts)
	queryString := fmt.Sprintf(dbCommand, planet.user, planet.dbID, scriptName)
	removeCommand := fmt.Sprintf("rm ~/sql/%s", scriptName)
	execCommand(queryString, planet, strucOut, opts)
	execCommand(removeCommand, planet, &placeholder, opts)
}
