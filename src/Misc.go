package main

import (
	"fmt"
	"os"
)

/**
################################################################################
							Miscallenous-Section
################################################################################
*/

/**
*	Prints the current Version of the goo application
 */
func printVersion() {
	fmt.Print("0.9")
}

/**
*	Prints the help dialog
 */
func printHelp() {
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
func throwErrOut(out []byte, err error) {
	fmt.Print(fmt.Sprint(err) + " output is: " + string(out) + "called from ErrOut. ")
	os.Stderr.WriteString(fmt.Sprint(err) + " output is: " + string(out) + "called from ErrOut. ")
	os.Exit(1)
}

/**
*	Formats and prints the given error.
 */
func throwErr(err error) {
	fmt.Print(fmt.Sprint(err) + " called from Err. ")
	os.Stderr.WriteString(fmt.Sprint(err) + "called from Err. ")
	os.Exit(1)
}
