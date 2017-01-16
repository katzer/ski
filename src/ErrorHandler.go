package main

import (
	"fmt"
	"log"
	"os"
)

/**
*	Formats and prints the given output and error.
 */
func throwErrOut(out []byte, err error) {
	errorString := fmt.Sprintf("%s output is: %s called from ErrOut.\n", fmt.Sprint(err), string(out))
	os.Stderr.WriteString(errorString)
	log.Output(2, `#################### ERROR ####################`)
	log.Output(2, errorString)
	log.Output(2, `#################### \ERROR ####################`)
	os.Exit(1)
}

/**
*	Formats and prints the given error.
 */
func throwErr(err error) {
	errorString := fmt.Sprintf("%s called from Err.\n", fmt.Sprint(err))
	os.Stderr.WriteString(errorString)
	log.Output(2, `#################### ERROR ####################`)
	log.Output(2, errorString)
	log.Output(2, `#################### \ERROR ####################`)
	os.Exit(1)
}

/**
*
 */
func throwErrExt(err error, addInf string) {
	errorString := fmt.Sprintf("%s AddInf: %s\n", fmt.Sprint(err), addInf)
	os.Stderr.WriteString(errorString)
	log.Output(2, `#################### ERROR ####################`)
	log.Output(2, errorString)
	log.Output(2, `#################### \ERROR ####################`)
	os.Exit(1)
}
