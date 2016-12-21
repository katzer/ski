package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/**
################################################################################
							ConsoleHandler-Section
################################################################################
*/

const planetLength int = 21

/**
*	Prints message to console with a newline.
 */
func println(msg string) {
	fmt.Println(msg)
}

/**
*	Prints message to console.
 */
func print(msg string) {
	fmt.Print(msg)
}

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
			println(buffer.String())
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
	println(buffer.String())
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

func formatAndPrint(toPrint []StructuredOuput, prettyFlag bool, scriptFlag bool, scriptPath string, command string) {
	if prettyFlag {
		printHeadline(scriptFlag, scriptPath, command, 26)
	}
	for i, planet := range toPrint {
		if !prettyFlag {
			println(planet.output)
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

func tablePrint(toFormat string, filePath string, templatePath string) {
	pys := getPyScript()
	pyScriptFile := os.Getenv("HOME") + "/tempTabFormat.py"
	err := ioutil.WriteFile(pyScriptFile, []byte(pys), 0644)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("python2", pyScriptFile, "/home/mrblati/workspace/goo/textfsm-master/examples/cisco_bgp_summary_template", "/home/mrblati/workspace/goo/textfsm-master/examples/cisco_bgp_summary_example")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" %q\n", out.String())

	err = os.Remove(pyScriptFile)
	if err != nil {
		log.Fatal(err)
	}
}
