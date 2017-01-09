package main

import (
	//"github.com/mgutz/ansi"
	"fmt"
	"gopkg.in/hypersleep/easyssh.v0"
	"os"
	"strings"
)

/**
################################################################################
								SSH-Section
################################################################################
*/

type SSHHandler struct {
}

/**
*	Executes a command on a remote ssh server
*	@params:
*		connDet: connection details in following form: user@hostname
*		cmd: command to be executed
 */
func execSSHCommand(connDet string, command string, strucOut *StructuredOuput, opts *Opts) {

	user := getUser(opts.currentDet)
	hostname := getHost(opts.currentDet)
	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:    os.Getenv("ORBIT_KEY"),
		Port:   "22",
	}
	var cmd string
	if opts.loadFlag {
		cmd = fmt.Sprintf(`sh -lc "echo -----APPPLANT-ORBIT----- && %s "`, command)
	} else {
		cmd = command
	}
	// Call Run method with command you want to run on remote server.
	out, err := ssh.Run(cmd)
	// Handle errors
	if err != nil {
		if opts.debugFlag {
			fmt.Println("#####SSH DEBUG#####")
			fmt.Printf("conndet: %s\n", connDet)
			fmt.Printf("user: %s\n", user)
			fmt.Printf("hostname: %s\n", hostname)
			fmt.Printf("orbit key: %s\n", os.Getenv("ORBIT_KEY"))
			fmt.Printf("command: %s\n", command)
			fmt.Printf("strucOut: %v\n", strucOut)
			fmt.Printf("opts:\n")
			fmt.Printf("prettyFlag: %t\n", opts.prettyFlag)
			fmt.Printf("scriptFlag: %t\n", opts.scriptFlag)
			fmt.Printf("typeFlag: %t\n", opts.typeFlag)
			fmt.Printf("debugFlag: %t\n", opts.debugFlag)
			fmt.Printf("loadFlag: %t\n", opts.loadFlag)
			fmt.Printf("helpFlag: %t\n", opts.helpFlag)
			fmt.Printf("versionFlag: %t\n", opts.versionFlag)
			fmt.Printf("tableFlag: %t\n", opts.tableFlag)
			fmt.Printf("scriptPath: %s\n", opts.scriptPath)
			fmt.Printf("command: %s\n", opts.command)
			fmt.Printf("planets: %v\n", opts.planets)
			fmt.Printf("planetsCount: %d\n", opts.planetsCount)
			fmt.Printf("currentDet: %s\n", opts.currentDet)
			fmt.Printf("currentDBDet: %s\n", opts.currentDBDet)

			fmt.Println("#####SSH DEBUG END#####")
		}
		throwErrExt(err, "called from exesSSHCommand ")
	} else {
		cleanedOut := out
		if opts.loadFlag {
			splitOut := strings.Split(out, "-----APPPLANT-ORBIT-----\n")
			cleanedOut = splitOut[len(splitOut)-1]
		}
		maxLength := getMaxLength(out)
		strucOut.output = cleanedOut
		strucOut.maxOutLength = maxLength
	}
}

/**
*	Uploads a file to the remote server
 */
func uploadFile(connDet string, opts *Opts) {
	user := getUser(opts.currentDet)
	hostname := getHost(opts.currentDet)

	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:    os.Getenv("ORBIT_KEY"),
		Port:   "22",
	}

	// Call Scp method with file you want to upload to remote server.
	err := ssh.Scp(opts.scriptPath)

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
func upAndExecSSHScript(connDet string, strucOut *StructuredOuput, opts *Opts) {
	uploadFile(connDet, opts)
	path := strings.Split(opts.scriptPath, "/")
	placeholder := StructuredOuput{}
	scriptName := path[len(path)-1]
	execSSHCommand(connDet, "chmod +x "+scriptName+" && "+"./"+scriptName, &placeholder, opts)
}
