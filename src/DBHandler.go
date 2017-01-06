package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/**
################################################################################
							DB-Handler-Section
################################################################################
*/

type DBHandler struct {
}

func execDBCommand(dbDet string, strucOut *StructuredOuput, opts *Opts) {
	tmpDBFile := os.Getenv("HOME") + "/orbit.sql"
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
	execSSHCommand(sshAddress, "mv ~/"+path[len(path)-1]+" ~/sql/"+path[len(path)-1], &placeholder, opts)
	queryString := ". profiles/" + username + ".prof && pqdb_sql.out -s " + dbID + " ~/sql/" + path[len(path)-1]
	//placeholder := StructuredOuput{}
	execSSHCommand(sshAddress, queryString, strucOut, opts)

}
