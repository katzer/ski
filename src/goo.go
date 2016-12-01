package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
    "gopkg.in/hypersleep/easyssh.v0"
)

/**
################################################################################
							Miscallenous-Section
################################################################################
*/

/**
*	Prints the current Version of the goo application
*/
func printVersion(){
	fmt.Println("0.0.1")
}

/**
*	Formats and prints the given output and error.
*/
func throwErr(out []byte, err error){
	fmt.Print(fmt.Sprint(err) + ": " + string(out) + "path is " + os.Getenv("PATH"))
	os.Stderr.WriteString(fmt.Sprint(err) + ": " + string(out))
	os.Exit(1)
}

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
func execCommand(connDet string, cmd string){
	user := getUser(connDet)
	hostname := getHost(connDet)

	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:  os.Getenv("ORBIT_KEY"),
		Port: "22",
	}

	// Call Run method with command you want to run on remote server.
	response, err := ssh.Run(cmd)
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println(response)
}

}


/**
################################################################################
						Information-Retrieval-Section
################################################################################
*/

/**
*	Splits the given connectiondetails and returns the hostname
*	@params:
*		connDet: Connection details in following form: user@hostname
*	@return: hostname
*/
func getHost(connDet string) string{
	toReturn := strings.Split(connDet,"@")
	return toReturn[1]
}

/**
*	Splits the given connectiondetails and returns the user
*	@params:
*		connDet: Connection details in following form: user@hostname
*	@return: user
*/
func getUser(connDet string) string{
	toReturn := strings.Split(connDet,"@")
	return toReturn[0]
}

/**
*	Returns the type of a given planet
*	@params:
*		id: The planets id
*	@return: The planets type
*/
func getType(id string) string{
	cmd 	 := exec.Command("ff","-t" ,id)
	out, err := cmd.CombinedOutput()
	if err != nil {
    	throwErr(out,err)
	}
	return strings.TrimSpace(string(out))
}

/**
*	Returns the connection details to a given planet
*	@params:
*		id: The planets id
*	@return: The connection details to the planet
*/
func getConnDet(id string) string{
	cmd 	 := exec.Command("ff",id)
	out, err := cmd.CombinedOutput()
	if err != nil {
    	throwErr(out,err)
	}
	return strings.TrimSpace(string(out))
}

/**
*	Extracts the desired argument from the arguments list.
*	@params:
*		args: Arguments to be searched in.
*		type: Type of desired Argument (command,id)
*	@return: The desired arguments
*/
func getArg(args []string, argType string) string{
	switch argType{
		case "command":
			var command string  = args[2]
			var cmdArgs []string
			if(len(args) > 3){
				cmdArgs = args[3:(len(args))]
				for _, argument := range cmdArgs {
					command += (" " + argument)
				}
			}
			return command
		case "id":
			return args[1]
		default:
			return "Unhandled Arg"
	}

}

/**
################################################################################
								Main-Section
################################################################################
*/

/**
*	Prints the help dialog
*/
func printHelp(){
	fmt.Println("usage: goo [options...] <planet>... <command>")
	fmt.Println("Options:")
	fmt.Println("-s, --script     Execute script and return result")
	fmt.Println("-p, --pretty     Pretty print output as a table")
	fmt.Println("-t, --type       Show type of planet")
	fmt.Println("-h, --help       This help text")
	fmt.Println("-v, --version    Show version number")
	os.Exit(1)
}

/**
*	Main function
*/
func main() {
	var args []string = os.Args
	prettyFlag := false
	typeFlag := false
	versionFlag := false
	scriptFlag := false

	fmt.Println(args)


	if(len(args) <3){
		printHelp()
	}
	for _, argument := range args {
		if(argument == "-h" || argument == "--help"){
			printHelp()
		}else if(argument == "-s" || argument == "--script"){
			scriptFlag = true
		}else if(argument == "-p" || argument == "--pretty"){
			prettyFlag = true
		}else if(argument == "-t" || argument == "--type"){
			typeFlag = true
		}else if(argument == "-v" || argument == "--version"){
			versionFlag = true
		}

	}

	if(versionFlag){
		printVersion()
	}
	_ = scriptFlag
	_ = prettyFlag
	_ = typeFlag
	switch getType(getArg(args,"id")) {
		case "server":
			var connDet string = getConnDet(getArg(args,"id"))
			var command string = getArg(args,"command")
			fmt.Println(connDet)
			fmt.Println("##########################")
			fmt.Println(command)
			fmt.Println("##########################")
			execCommand(connDet,command)
		case "db":
			fmt.Println("This Type of Connection is not yet supported.")
			os.Exit(1)
		case "web":
			fmt.Println("This Type of Connection is not supported.")
			os.Exit(1)
		default:
			fmt.Println(getType(getArg(args,"id")))
			fmt.Println("###")
			fmt.Println(getArg(args,"id"))

	}
}
