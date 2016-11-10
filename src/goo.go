package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	//"strconv"
)

func main() {
	if(len(os.Args) == 1){
		fmt.Println("usage: goo <Server-ID> <Command> [Arguments]")
		os.Exit(0)
	}
	var cmdArgs []string
	idString := os.Args[1]
	if(len(os.Args) >= 3){
		cmdArgs = os.Args[2:(len(os.Args))]
	}

	//cmdArgs := os.Args[3:len(os.Args)]
	cmd 	 := exec.Command("ff",idString)
	out, err := cmd.CombinedOutput()

	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(out))
    		os.Exit(1)
	}
	connectionData := string(out)

	cmd 	 = exec.Command("ff","-t" ,idString)
	out, err = cmd.CombinedOutput()
	if err != nil {
    		fmt.Println(fmt.Sprint(err) + ": " + string(out))
    		os.Exit(1)
	}
	planetType := strings.TrimSpace(string(out))
	fmt.Println(planetType)
	switch planetType {
		case "server":
			cmd 	 := exec.Command("ssh",connectionData)
			out, err := cmd.CombinedOutput()
			fmt.Println(string(out))
			if err != nil {
    				fmt.Println(fmt.Sprint(err) + ": " + string(out))
    				os.Exit(1)
			}
			if len(os.Args) >= 3 {
				for _, command := range cmdArgs {
					cmd 	 := exec.Command("ssh",connectionData,command)
					out, err := cmd.CombinedOutput()
					fmt.Println(string(out))
					if err != nil {
    						fmt.Println(fmt.Sprint(err) + ": " + string(out))
    						os.Exit(1)
					}
				}
			}

		case "db":
			fmt.Println("This Type of Connection is not yet supportet")
			os.Exit(1)
		case "web":
			fmt.Println("This Type of Connection is not yet supportet")
			os.Exit(1)
		default:
			fmt.Println("Sag doch irgendetwas")

	}

}
