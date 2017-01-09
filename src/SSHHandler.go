package main

import (
	//"github.com/mgutz/ansi"
	"fmt"
	"gopkg.in/hypersleep/easyssh.v0"
	"os"
	"strings"
)

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
		cmd = "sh -lc \"echo -----APPPLANT-ORBIT----- && " + command + " \""
	} else {
		cmd = command
	}
	// Call Run method with command you want to run on remote server.
	out, err := ssh.Run(cmd)
	// Handle errors
	if err != nil {
		if opts.debugFlag {
			fmt.Println("#####SSH DEBUG#####")
			fmt.Println("conndet:")
			fmt.Println(connDet)
			fmt.Println("user:")
			fmt.Println(user)
			fmt.Println("hostname:")
			fmt.Println(hostname)
			fmt.Println("orbit key:")
			fmt.Println(os.Getenv("ORBIT_KEY"))
			fmt.Println("command:")
			fmt.Println(command)
			fmt.Println("strucOut:")
			fmt.Println(strucOut)
			fmt.Println("opts:")
			fmt.Print("prettyFlag: ")
			fmt.Println(opts.prettyFlag)
			fmt.Print("scriptFlag: ")
			fmt.Println(opts.scriptFlag)
			fmt.Print("typeFlag: ")
			fmt.Println(opts.typeFlag)
			fmt.Print("debugFlag: ")
			fmt.Println(opts.debugFlag)
			fmt.Print("loadFlag: ")
			fmt.Println(opts.loadFlag)
			fmt.Print("helpFlag: ")
			fmt.Println(opts.helpFlag)
			fmt.Print("versionFlag: ")
			fmt.Println(opts.versionFlag)
			fmt.Print("tableFlag: ")
			fmt.Println(opts.tableFlag)
			fmt.Print("scriptPath: ")
			fmt.Println(opts.scriptPath)
			fmt.Print("command: ")
			fmt.Println(opts.command)
			fmt.Print("planets: ")
			fmt.Println(opts.planets)
			fmt.Print("planetsCount: ")
			fmt.Println(opts.planetsCount)
			fmt.Print("currentDet: ")
			fmt.Println(opts.currentDet)
			fmt.Print("currentDBDet: ")
			fmt.Println(opts.currentDBDet)

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
