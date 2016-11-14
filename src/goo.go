package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	//"strconv"
)

func throwErr(out []byte, err error){
	fmt.Print(fmt.Sprint(err) + ": " + string(out) + "path is " + os.Getenv("PATH"))
	os.Stderr.WriteString(fmt.Sprint(err) + ": " + string(out))
	os.Exit(1)
}

func execCommand(connectionData string, command string){
	cmd 	 := exec.Command("ssh",connectionData, command)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		throwErr(out,err)
	}
}

func getType(args []string) string{
	cmd 	 := exec.Command("ff","-t" ,args[1])
	out, err := cmd.CombinedOutput()
	if err != nil {
    	throwErr(out,err)
	}
	return strings.TrimSpace(string(out))
}

func getConnDet(args []string) string{
	cmd 	 := exec.Command("ff",args[1])
	out, err := cmd.CombinedOutput()
	if err != nil {
    	throwErr(out,err)
	}
	return strings.TrimSpace(string(out))
}

func getCommand(args []string) string{
	var command string  = args[2]
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
	if (len(args) <=2) || args[1] == "-h"{
		fmt.Println("usage: goo <Server-ID> <Command> [Arguments]")
		os.Exit(0)
	}
	switch getType(args) {
		case "server":
			var connDet string = getConnDet(args)
			var command string = getCommand(args)
			execCommand(connDet,command)
		case "db":
			fmt.Println("This Type of Connection is not yet supported.")
			os.Exit(1)
		default:
			fmt.Println("This Type of Connection is not supported.")
			os.Exit(1)
	}
}
