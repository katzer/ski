package main

import (
	"fmt"
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"os"
	"path"
)

func execDBCommand(dbID string, user string, hostname string, strucOut *StructuredOuput, opts *Opts) {
	tmpDBFile := path.Join(os.Getenv("ORBIT_HOME"), "scripts", "orbit.sql")
	err := ioutil.WriteFile(tmpDBFile, []byte(opts.command), 0644)
	if err != nil {
		log.Fatalf("writing temporary sql script failed : %v", err)
	}
	opts.scriptName = "orbit.sql"
	execDBScript(dbID, user, hostname, strucOut, opts)
	err = os.Remove(tmpDBFile)
	if err != nil {
		log.Fatal(err)
	}
}

func execDBScript(dbID string, user string, hostname string, strucOut *StructuredOuput, opts *Opts) {
	const dbCommand = ". profiles/%s.prof && pqdb_sql.out -s %s ~/sql/%s"
	uploadFile(user, hostname, opts)
	placeholder := StructuredOuput{}
	scriptName := opts.scriptName
	command := fmt.Sprintf("mv ~/%s ~/sql/%s", scriptName, scriptName)
	execCommand(user, hostname, command, &placeholder, opts)
	queryString := fmt.Sprintf(dbCommand, user, dbID, scriptName)
	removeCommand := fmt.Sprintf("rm ~/sql/%s", scriptName)
	execCommand(user, hostname, queryString, strucOut, opts)
	execCommand(user, hostname, removeCommand, &placeholder, opts)
}
