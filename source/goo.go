package main

import (
	"fmt"
	"os"
	"os/exec"
	//"strconv"
)

func main() {
	if(len(os.Args) == 1){
		fmt.Println("usage: goo <Server-ID> <Command> [Arguments]")
		os.Exit(0)
	}
	idString := os.Args[1]
	cmdName := os.Args[2]
	//cmdArgs := os.Args[3:len(os.Args)]
	cmd 	 := exec.Command("ff",idString)
	out, err := cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(out))
    		return
	}
	connectionData := string(out)
	fmt.Println("The id is", idString)
	fmt.Println("The command is", cmdName)
	fmt.Println("The output is", connectionData)

	cmd 	 = exec.Command("ff","-t" ,idString)
	out, err = cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(out))
    		return
	}
	planetType := string(out)
	switch planetType{
		case "server":
			//do server
		case "db":
			//do db
		case "web":
			//do web
		default:

	}

	fmt.Println("Type is", planetType)
}
