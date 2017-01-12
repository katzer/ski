package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func execDBCommand(dbID string, user string, hostname string, strucOut *StructuredOuput, opts *Opts) {
	tmpDBFile := fmt.Sprintf("%s/orbit.sql", os.Getenv("HOME"))
	err := ioutil.WriteFile(tmpDBFile, []byte(opts.command), 0644)
	if err != nil {
		fmt.Println("writefile failed!")
		os.Exit(1)
	}
	opts.scriptPath = tmpDBFile
	execDBScript(dbID, user, hostname, strucOut, opts)
	err = os.Remove(tmpDBFile)
	if err != nil {
		log.Fatal(err)
	}
}

func execDBScript(dbID string, user string, hostname string, strucOut *StructuredOuput, opts *Opts) {
	uploadFile(user, hostname, opts)
	path := strings.Split(opts.scriptPath, "/")
	placeholder := StructuredOuput{}
	scriptName := path[len(path)-1]
	command := fmt.Sprintf("mv ~/%s ~/sql/%s", scriptName, scriptName)
	execCommand(user, hostname, command, &placeholder, opts)
	queryString := fmt.Sprintf(". profiles/%s.prof && pqdb_sql.out -s %s ~/sql/%s", user, dbID, scriptName)
	execCommand(user, hostname, queryString, strucOut, opts)

}
