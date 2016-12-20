package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	//"os"
	//"strconv"
	//"sync"
	"os"
)

/**
################################################################################
							DB-Handler-Section
################################################################################
*/

func execDBCommand(dbDet string, command string, wg *sync.WaitGroup, wait bool, strucOut *StructuredOuput, loadFlag bool) {
	dbID, sshAddress := procDBDets(dbDet)
	err := ioutil.WriteFile("/tmp/dbDAT", []byte(command), 0644)
	if err != nil {
		println("writefile failed!")
		os.Exit(1)
	}
	uploadFile(sshAddress, "/tmp/dbDAT")
	queryString := "pqdb " + dbID + " ~/dbDAT"
	execSSHCommand(sshAddress, queryString, wg, true, strucOut, loadFlag)
}

func upAndExecDBScript(dbDet string, scriptPath string, wg *sync.WaitGroup, strucOut *StructuredOuput, loadFlag bool) {
	dbID, sshAddress := procDBDets(dbDet)
	uploadFile(sshAddress, scriptPath)
	path := strings.Split(scriptPath, "/")
	queryString := "pqdb " + dbID + " ~/" + path[len(path)-1]
	//placeholder := StructuredOuput{}
	execSSHCommand(sshAddress, queryString, wg, true, strucOut, loadFlag)
}
