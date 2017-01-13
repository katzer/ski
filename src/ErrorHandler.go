package main

import (
	"fmt"
	"os"
)

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
	os.Stderr.WriteString(fmt.Sprint(err) + " called from Err. ")
	os.Exit(1)
}

/**
*
 */
func throwErrExt(err error, addInf string) {
	fmt.Print(fmt.Sprint(err) + " AddInf: " + addInf)
	os.Stderr.WriteString(fmt.Sprint(err) + " AddInf: " + addInf)
	os.Exit(1)
}
