package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type DBHandler struct {
}

func execDBCommand(dbDet string, strucOut *StructuredOuput, opts *Opts) {
	tmpDBFile := fmt.Sprintf("%s/orbit.sql", os.Getenv("HOME"))
	err := ioutil.WriteFile(tmpDBFile, []byte(opts.command), 0644)
	if err != nil {
		fmt.Println("writefile failed!")
		os.Exit(1)
	}
	opts.scriptPath = tmpDBFile
	upAndExecDBScript(dbDet, strucOut, opts)
	err = os.Remove(tmpDBFile)
	if err != nil {
		log.Fatal(err)
	}
}

func upAndExecDBScript(dbDet string, strucOut *StructuredOuput, opts *Opts) {
	dbID, sshAddress := procDBDets(dbDet)
	username := getUser(sshAddress)
	uploadFile(sshAddress, opts)
	path := strings.Split(opts.scriptPath, "/")
	placeholder := StructuredOuput{}
	scriptName := path[len(path)-1]
	command := fmt.Sprintf("mv ~/%s ~/sql/%s", scriptName, scriptName)
	execSSHCommand(sshAddress, command, &placeholder, opts)
	queryString := fmt.Sprintf(". profiles/%s.prof && pqdb_sql.out -s %s ~/sql/%s", username, dbID, scriptName)
	execSSHCommand(sshAddress, queryString, strucOut, opts)

}
