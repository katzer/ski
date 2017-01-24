package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func execDBCommand(dbID string, user string, hostname string, strucOut *StructuredOuput, opts *Opts) {
	tmpDBFile := ""
	if runtime.GOOS == "windows" {
		tmpDBFile = path.Join(os.Getenv("TEMP"), "orbit.sql")
	} else {
		tmpDBFile = path.Join(os.Getenv("HOME"), "orbit.sql")
	}
	err := ioutil.WriteFile(tmpDBFile, []byte(opts.command), 0644)
	if err != nil {
		fmt.Println("writing temporary sql script failed")
		log.Fatal(err)
	}
	opts.scriptPath = tmpDBFile
	execDBScript(dbID, user, hostname, strucOut, opts)
	err = os.Remove(tmpDBFile)
	if err != nil {
		log.Fatal(err)
	}
}

func execDBScript(dbID string, user string, hostname string, strucOut *StructuredOuput, opts *Opts) {
	const dbCommand = ". profiles/%s.prof && pqdb_sql.out -s %s ~/sql/%s"
	uploadFile(user, hostname, opts)
	path := strings.Split(opts.scriptPath, "/")
	if runtime.GOOS == "windows" {
		path = strings.Split(opts.scriptPath, "\\")
	} else {
		path = strings.Split(opts.scriptPath, "/")
	}
	placeholder := StructuredOuput{}
	scriptName := path[len(path)-1]
	command := fmt.Sprintf("mv ~/%s ~/sql/%s", scriptName, scriptName)
	removeCommand := fmt.Sprintf("rm ~/sql/%s", scriptName)
	execCommand(user, hostname, command, &placeholder, opts)
	queryString := fmt.Sprintf(dbCommand, user, dbID, scriptName)
	execCommand(user, hostname, queryString, strucOut, opts)
	execCommand(user, hostname, removeCommand, &placeholder, opts)
}
