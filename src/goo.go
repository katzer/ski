package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
    "gopkg.in/hypersleep/easyssh.v0"
    "strconv"
    "sync"
    "github.com/mgutz/ansi"
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
*	Prints the help dialog
*/
func printHelp(){
	fmt.Println("usage: goo [options...] <planet>... -c=\"<command>\"")
	fmt.Println("Options:")
	fmt.Println("-s=\"<path/to/script>\", --script=\"<path/to/script>\"  Execute script and return result")
	fmt.Println("-p, --pretty     Pretty print output as a table")
	fmt.Println("-t, --type       Show type of planet")
	fmt.Println("-h, --help       This help text")
	fmt.Println("-v, --version    Show version number")
	fmt.Println("-d, --debug	  Show extended debug informations")
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
func execCommand(connDet string, cmd string, wg *sync.WaitGroup , wait bool, prettyFlag bool){

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
		if(prettyFlag){
			fmt.Println(ansi.Color("################################################################################","blue"))
			fmt.Println(ansi.Color("Executing command ","241") + ansi.Color(cmd,"white+hu") + ansi.Color(" on ","241") + ansi.Color(connDet,"white+hu"))
		}else{
			fmt.Println("Executing command " + cmd + " on " + connDet)
		}
		fmt.Println("")
		fmt.Println(out)
	}
	if (wait){
		wg.Done()
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
*	Uploads and executes a script on a given planet
*	@params:
*		connDet: 	Connection details to planet
*		scriptPath: Path to script
*/
func upAndExecScript(connDet string, scriptPath string, wg *sync.WaitGroup, prettyFlag bool){
	uploadFile(connDet,scriptPath)
	path := strings.Split(scriptPath,"/")
	execCommand(connDet,"chmod +x " + path[len(path)-1],wg,false,prettyFlag)
	execCommand(connDet,"./" + path[len(path)-1],wg,false,prettyFlag)
	wg.Done()
}


/**
################################################################################
						Information-Retrieval-Section
################################################################################
*/

/**
*	Returns the contents of args in following order:
*	prettyprint flag
*	script flag
*	script path
*	command
*	planets
*/
func procArgs(args []string) (bool, bool, string, string, []string, bool, bool){
	prettyFlag := false
	scriptFlag := false
	typeFlag := false
	debugFlag := false
	var scriptPath string = ""
	var command string = ""
	planets := make([]string,0)

	cleanArgs := args[1:]



	for _, argument := range cleanArgs {
		if(strings.HasPrefix(argument,"-h") || strings.HasPrefix(argument,"--help")){
			printHelp()
			os.Exit(0)
		}else if(strings.HasPrefix(argument,"-p") || strings.HasPrefix(argument,"--pretty")){
			prettyFlag = true
		}else if(strings.HasPrefix(argument,"-t") || strings.HasPrefix(argument,"--type")){
			typeFlag = true
		}else if(strings.HasPrefix(argument,"-d") || strings.HasPrefix(argument,"--debug")){
			debugFlag = true
		}else if(strings.HasPrefix(argument,"-v") || strings.HasPrefix(argument,"--version")){
			printVersion()
			os.Exit(0)
		}else if(strings.HasPrefix(argument,"-c") || strings.HasPrefix(argument,"--command")){
			// TODO what if theres a = in the command itself?
			command = strings.TrimSuffix(strings.TrimPrefix(strings.Split(argument,"=")[1],"\""),"\"")
		}else if(strings.HasPrefix(argument,"-s") || strings.HasPrefix(argument,"--script")){
			scriptFlag = true
			scriptPath = strings.Split(argument,"=")[1]
		}else{
			planets = append(planets,argument)
		}
	}
	if(len(args) <3){
		printHelp()
		os.Exit(0)
	}

	_ = prettyFlag

	return prettyFlag,scriptFlag,scriptPath,command,planets,debugFlag,typeFlag
}


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
*					DEPRECATED
*
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
*	Main function
*/
func main() {

	var args []string = os.Args

	prettyFlag,scriptFlag,scriptPath,command,planets,debugFlag,typeFlag := procArgs(args)

	_ = prettyFlag
	if(debugFlag){
		fmt.Println(args)
		fmt.Println("prettyflag " + strconv.FormatBool(prettyFlag))
		fmt.Println("scriptflag " + strconv.FormatBool(scriptFlag))
		fmt.Println("scriptpath " + scriptPath)
		fmt.Println("command " + command)
		for _, planet := range planets {
			fmt.Println("planet " + planet)
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(planets))
	for _, planet := range planets {
		if(typeFlag){
			fmt.Println("The type of " + planet + " is " + getType(planet))
		}
		switch getType(planet) {
			case "server":
				var connDet string = getConnDet(planet)
				if(scriptFlag){
					go upAndExecScript(connDet,scriptPath,&wg,prettyFlag)
				}else{
					go execCommand(connDet,command,&wg,true,prettyFlag)
				}
			case "db":
				fmt.Fprintf(os.Stderr, "This Type of Connection is not yet supported.")
				os.Exit(1)
			case "web":
				fmt.Fprintf(os.Stderr, "This Type of Connection is not supported.")
				os.Exit(1)
			default:
				wg.Done()
		}
	}
	wg.Wait()
}
