package main

import (
	"fmt"
	"strconv"
	"bytes"
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
	indentCheck := 80 - indent
	if(!exceptFirst){
		for i := 0; i < indent; i++ {
			buffer.WriteByte(32)
		}
		indentCheck = 80
	}
	for _, char := range charString {
		if(buffer.Len() >= indentCheck || char == 10){
			indentCheck = 80
			println(buffer.String())
			buffer.Reset()
			for i := 0; i < indent; i++ {
				buffer.WriteByte(32)
			}
		}
		if(char != 10){
			toAppend = fmt.Sprintf("%c", char)
			buffer.WriteString(toAppend)
		}
	}
	println(buffer.String())
}

func printHeadline(scriptFlag bool, scriptPath string, command string, indent int){
	print("NR   PLANET               ")
	if(scriptFlag){
		printIndented(scriptPath,indent,true)
	}else{
		printIndented(command,indent,true)
	}
	fmt.Println("================================================================================")
}

func printWhite(length int) {
	for i := 0; i < length; i++{
		print(" ")
	}
}

func formatAndPrint(toPrint []StructuredOuput, prettyFlag bool, scriptFlag bool, scriptPath string, command string) {
	if(prettyFlag){
		printHeadline(scriptFlag, scriptPath, command, 26)
	}
	for i, planet := range toPrint {
		if(!prettyFlag){
			println(planet.planet)
			println(planet.output)
		}else{
			print(strconv.Itoa(i)+ "")
			if(i/10 < 1){
				print(" ")
			}
			if(i/100 < 1){
				print(" ")
			}
			if(i/1000 < 1){
				print(" ")
			}
			print(" ")
			print(planet.planet)
			printWhite(planetLength - len(planet.planet))
			printIndented(planet.output,26,true)
		}
	}
}
