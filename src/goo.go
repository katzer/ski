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
	if(len(os.Args) == 3){
		//cmdName := os.Args[2]
	}

	//cmdArgs := os.Args[3:len(os.Args)]
	cmd 	 := exec.Command("ff",idString)
	out, err := cmd.CombinedOutput()

	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(out))
    		return
	}
	connectionData := string(out)

	fmt.Println(connectionData)

	cmd 	 = exec.Command("ff","-t" ,idString)
	out, err = cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(out))
    		return
	}
	planetType := string(out)
	fmt.Println("Type is", planetType)
	switch planetType{
		case "server":
			//cmd 	 := exec.Command("ssh",idString)
			//out, err := cmd.CombinedOutput()
		case "db":
			//cmd 	 := exec.Command("ssh",idString)
			//out, err := cmd.CombinedOutput()
		case "web":
			//cmd 	 := exec.Command("ssh",idString)
			//out, err := cmd.CombinedOutput()
		default:

	}

}
