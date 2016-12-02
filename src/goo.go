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
func throwErrOut(out []byte, err error){
	fmt.Print(fmt.Sprint(err) + " output is: " + string(out) + "called from ErrOut. ")
	os.Stderr.WriteString(fmt.Sprint(err) + " output is: " + string(out) + "called from ErrOut. ")
	os.Exit(1)
}

/**
*	Formats and prints the given error.
*/
func throwErr(err error){
	fmt.Print(fmt.Sprint(err)  + " called from Err. ")
	os.Stderr.WriteString(fmt.Sprint(err) + "called from Err. ")
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
	out, err := ssh.Run(cmd)
	// Handle errors
	if err != nil {
		throwErr(err)
	} else {
		fmt.Println(out)
	}
}

/**
*	Uploads a file to the remote server
*/
func uploadFile(connDet string, path string){
	user := getUser(connDet)
	hostname := getHost(connDet)

	ssh := &easyssh.MakeConfig{
		User:   user,
		Server: hostname,
		Key:  os.Getenv("ORBIT_KEY"),
		Port: "22",
	}

	// Call Scp method with file you want to upload to remote server.
	err := ssh.Scp(path)

	// Handle errors
	if err != nil {
		throwErr(err)
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
    	throwErrOut(out,err)
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
    	throwErrOut(out,err)
	}
	return strings.TrimSpace(string(out))
}

/**
*	Extracts the desired argument from the arguments list.
*	@params:
*		args: Arguments to be searched in.
*		type: Type of desired Argument (command,id)
*		position: starting position of desired argument
*	@return: The desired arguments
*/
func getArg(args []string, argType string, position int) string{
	switch argType{
		case "command":
			var command string  = args[position]
			var cmdArgs []string
			if(len(args) > (position+1)){
				cmdArgs = args[(position+1):(len(args))]
				for _, argument := range cmdArgs {
					command += (" " + argument)
				}
			}
			return command
		default:
			return args[position]
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
	fmt.Println("-s <path/to/script>, --script  <path/to/script>   Execute script and return result")
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
	commandsPosition := 2
	planetPosition := 1
	scriptPosition := 0
	index := 0
	_ = scriptPosition
	fmt.Println(args)


	if(len(args) <3){
		printHelp()
	}
	for _, argument := range args {
		if(argument == "-h" || argument == "--help"){
			printHelp()
		}else if(argument == "-s" || argument == "--script"){
			scriptFlag = true
			planetPosition += 2
			scriptPosition = index + 1
		}else if(argument == "-p" || argument == "--pretty"){
			prettyFlag = true
			commandsPosition ++
			planetPosition ++
		}else if(argument == "-t" || argument == "--type"){
			typeFlag = true
			commandsPosition ++
			planetPosition ++
		}else if(argument == "-v" || argument == "--version"){
			versionFlag = true
			commandsPosition ++
			planetPosition ++
		}
		index ++
	}

	if(versionFlag){
		printVersion()
	}
	if(typeFlag){
		fmt.Println(getType(getArg(args,"id",planetPosition)))
	}
	_ = prettyFlag
	switch getType(getArg(args,"id",planetPosition)) {
		case "server":
			var connDet string = getConnDet(getArg(args,"id",planetPosition))
			if(scriptFlag){
				scriptFile := getArg(args,"scriptFile",scriptPosition)
				uploadFile(connDet,scriptFile)
				path := strings.Split(scriptFile,"/")
				execCommand(connDet,"chmod +x " + path[len(path)-1])
				execCommand(connDet,"./" + path[len(path)-1])
			}else{
				command := getArg(args,"command",commandsPosition)
				execCommand(connDet,command)
			}
		case "db":
			fmt.Println("This Type of Connection is not yet supported.")
			os.Exit(1)
		case "web":
			fmt.Println("This Type of Connection is not supported.")
			os.Exit(1)
		default:
			fmt.Println(getType(getArg(args,"id",planetPosition)))
			fmt.Println("###")
			fmt.Println(getArg(args,"id",planetPosition))

	}
}
