package main

import log "github.com/Sirupsen/logrus"

func throwErrOut(out []byte, err error) {
	log.Fatalf("%s output is: %s called from ErrOut.\n", err, out)
}

func throwErr(err error) {
	log.Fatal(err)
}

func throwErrExt(err error, addInf string) {
	log.Fatalf("%s\nAdditional info: %s\n", err, addInf)
}
