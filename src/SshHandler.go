package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"gopkg.in/hypersleep/easyssh.v0"
	"os"
	"strings"
	"sync"
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
func execCommand(connDet string, cmd string, wg *sync.WaitGroup, wait bool, prettyFlag bool) {

	user := getUser(connDet)
	hostname := getHost(connDet)

	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:    os.Getenv("ORBIT_KEY"),
		Port:   "22",
	}
	// Call Run method with command you want to run on remote server.
	out, err := ssh.Run(cmd)
	// Handle errors
	if err != nil {
		throwErr(err)
	} else {
		if prettyFlag {
			fmt.Println(ansi.Color("################################################################################", "blue"))
			fmt.Println(ansi.Color("Executing command ", "241") + ansi.Color(cmd, "white+hu") + ansi.Color(" on ", "241") + ansi.Color(connDet, "white+hu"))
		} else {
			fmt.Println("Executing command " + cmd + " on " + connDet)
		}
		fmt.Println("")
		fmt.Println(out)
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
func upAndExecScript(connDet string, scriptPath string, wg *sync.WaitGroup, prettyFlag bool) {
	uploadFile(connDet, scriptPath)
	path := strings.Split(scriptPath, "/")
	execCommand(connDet, "chmod +x "+path[len(path)-1], wg, false, prettyFlag)
	execCommand(connDet, "./"+path[len(path)-1], wg, false, prettyFlag)
	wg.Done()
}

