package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	//"strconv"
)



func execCommand(connectionData string, command string){
	cmd 	 := exec.Command("ssh",connectionData, command)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + string(out))
			os.Exit(1)
	}
}

func getType string(args []string){
	cmd 	 := exec.Command("ff","-t" ,args[1])
	out, err := cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(out))
    		os.Exit(1)
	}
	return strings.TrimSpace(string(out))
}

func getConnDet string(args []string){
	cmd 	 := exec.Command("ff",args[1])
	out, err := cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(out))
    		os.Exit(1)
	}
	return strings.TrimSpace(string(out))
}

func getCommand string(args []string){
	var string command = args[2]
	var cmdArgs []string
	if(len(args) >= 3){
		cmdArgs = args[2:(len(args))]
	}
	if len(args) >= 3 {
		for _, argument := range cmdArgs {
			command += (" " + argument)
		}
	}
	return command
}

func main() {
	var args []string = os.Args
	if len.(args) <=2 || args[1] == "-h"{
		fmt.Println("usage: goo <Server-ID> <Command> [Arguments]")
		os.Exit(0)
	}
	switch getType(args) {
		case "server":
			var connDet string = getConnDet(args)
			var command string = getCommand(args)
			execCommand(connDet,command)
		case "db":
			fmt.Println("This type of connection ist not yet supported.")
			os.Exit(1)
		default:
			fmt.Println("This type connection is not supported.")
			os.Exit(1)
	}
}
