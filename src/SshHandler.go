package main

import (
	//"github.com/mgutz/ansi"
	"strings"
	"sync"
	"gopkg.in/hypersleep/easyssh.v0"
	"os"
)

/**
################################################################################
								SSH-Section
################################################################################
*/

/**
*	Executes a command on a remote ssh server
*	@params:
*		connDet: connection details in following form: user@hostname
*		cmd: command to be executed
 */
func execCommand(connDet string, cmd string, wg *sync.WaitGroup, wait bool, strucOut *StructuredOuput, loadFlag bool) {

	user := getUser(connDet)
	hostname := getHost(connDet)
	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:    os.Getenv("ORBIT_KEY"),
		Port:   "22",
	}
	if(loadFlag){
		cmd = "sh -lc " + cmd
	}
	// Call Run method with command you want to run on remote server.
	out, err := ssh.Run(cmd)
	// Handle errors
	if err != nil {
		throwErr(err)
	} else {
		maxLength := getMaxLength(out)
		strucOut.output = out
		strucOut.maxOutLength = maxLength
	}
	if wait {
		wg.Done()
	}
}

/**
*	Uploads a file to the remote server
 */
func uploadFile(connDet string, path string) {
	user := getUser(connDet)
	hostname := getHost(connDet)

	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:    os.Getenv("ORBIT_KEY"),
		Port:   "22",
	}

	// Call Scp method with file you want to upload to remote server.
	err := ssh.Scp(path)

	// Handle errors
	if err != nil {
		throwErr(err)
	}
}

/**
*	Uploads and executes a script on a given planet
*	@params:
*		connDet: 	Connection details to planet
*		scriptPath: Path to script
 */
func upAndExecScript(connDet string, scriptPath string, wg *sync.WaitGroup, strucOut *StructuredOuput, loadFlag bool) {
	uploadFile(connDet, scriptPath)
	path := strings.Split(scriptPath, "/")
	placeholder := StructuredOuput{}
	execCommand(connDet, "chmod +x "+path[len(path)-1], wg, false,&placeholder, false)
	execCommand(connDet, "./"+path[len(path)-1], wg, false,strucOut, loadFlag)
	wg.Done()
}

