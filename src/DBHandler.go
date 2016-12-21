package main

import (
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

func execDBCommand(dbDet string, command string, wg *sync.WaitGroup, strucOut *StructuredOuput, loadFlag bool) {
	tmpDBFile := os.Getenv("HOME") + "/orbit.sql"
	err := ioutil.WriteFile(tmpDBFile, []byte(command), 0644)
	if err != nil {
		println("writefile failed!")
		os.Exit(1)
	}
	upAndExecDBScript(dbDet, tmpDBFile, wg, strucOut, loadFlag)
	err = os.Remove(tmpDBFile)
	if err != nil {
		log.Fatal(err)
	}
}

func upAndExecDBScript(dbDet string, scriptPath string, wg *sync.WaitGroup, strucOut *StructuredOuput, loadFlag bool) {
	dbID, sshAddress := procDBDets(dbDet)
	username := getUser(sshAddress)
	uploadFile(sshAddress, scriptPath)
	path := strings.Split(scriptPath, "/")
	placeholder := StructuredOuput{}
	execSSHCommand(sshAddress, "mv ~/"+path[len(path)-1]+" ~/sql/"+path[len(path)-1], wg, false, placeholder, loadFlag)
	queryString := ". profiles/" + username + ".prof && pqdb_sql.out -s " + dbID + " ~/sql/" + path[len(path)-1]
	//placeholder := StructuredOuput{}
	execSSHCommand(sshAddress, queryString, wg, true, strucOut, loadFlag)

}
