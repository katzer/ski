package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		cmdOut []byte
		err    error
	)
	idString := os.Args[1]
	cmdName := os.Args[2]
	cmdArgs := os.Args[3:len(os.Args)]
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error the command: ", err)
		//os.Exit(1)
	}
	sha := string(cmdOut)
	fmt.Println("The id is", idString)
	fmt.Println("The command is", cmdName)
	fmt.Println("The output is", sha)
}