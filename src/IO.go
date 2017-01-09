package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const planetLength int = 21

/**
*	Prints the given String indented by the given spaces.
 */
func printIndented(msg string, indent int, exceptFirst bool) {
	charString := []rune(msg)
	var buffer bytes.Buffer
	_ = exceptFirst
	toAppend := ""
	if !exceptFirst {
		for i := 0; i < indent; i++ {
			buffer.WriteByte(32)
		}
	}
	for _, char := range charString {
		if char == 10 {
			fmt.Println(buffer.String())
			buffer.Reset()
			for i := 0; i < indent; i++ {
				buffer.WriteByte(32)
			}
		}
		if char != 10 {
			toAppend = fmt.Sprintf("%c", char)
			buffer.WriteString(toAppend)
		}
	}
	fmt.Println(buffer.String())
}

func printHeadline(scriptFlag bool, scriptPath string, command string, indent int) {
	print("NR   PLANET               ")
	if scriptFlag {
		printIndented(scriptPath, indent, true)
	} else {
		printIndented(command, indent, true)
	}
	fmt.Println("================================================================================")
}

func printWhite(length int) {
	for i := 0; i < length; i++ {
		print(" ")
	}
}

func formatAndPrint(toPrint []StructuredOuput, opts *Opts) {
	if opts.prettyFlag {
		printHeadline(opts.scriptFlag, opts.scriptPath, opts.command, 26)
	}
	for i, planet := range toPrint {
		if !opts.prettyFlag {
			print(planet.output)
		} else {
			print(strconv.Itoa(i) + "")
			if i/10 < 1 {
				print(" ")
			}
			if i/100 < 1 {
				print(" ")
			}
			if i/1000 < 1 {
				print(" ")
			}
			print(" ")
			print(planet.planet)
			printWhite(planetLength - len(planet.planet))
			printIndented(planet.output, 26, true)
		}
	}
}

func tablePrint(templatePath string, filePath string) {
	pys := getPyScript()
	pyScriptFile := ""
	if runtime.GOOS == "windows" {
		pyScriptFile = os.Getenv("TEMP") + "\\tempTabFormat.py"
	} else {
		pyScriptFile = os.Getenv("HOME") + "/tempTabFormat.py"
	}
	err := ioutil.WriteFile(pyScriptFile, []byte(pys), 0644)
	if err != nil {
		fmt.Println("writing pyscript failed")
		log.Fatal(err)
	}
	cmd := exec.Command("python2", pyScriptFile, templatePath, filePath)
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("executing pyscript failed")
		log.Fatal(err)
	}
	formattedString := strings.Split(out.String(), "FSM Table:\n")[1]
	fmt.Printf(" %q\n", formattedString)

	err = os.Remove(pyScriptFile)
	if err != nil {
		fmt.Println("removing pyscript failed")
		log.Fatal(err)
	}
}
