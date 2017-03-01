package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

var toUnMarshall string
var data []byte
var absPath string

func setup() {
	toUnMarshall = "job.js"
	json := `{
  "debug"      : true,
  "help"       : true,
  "load"       : true,
  "version"    : false,
  "command"    : "whatever",
  "scriptName" : "doink",
  "template"   : "not_a_good_one",
  "planets"    : [ "1","2","3"]
}
`
	data = []byte(json)
	absPath = path.Join(os.TempDir(), toUnMarshall)
	err := ioutil.WriteFile(absPath, data, 0644)
	if err != nil {
		panic(err)
	}
}

func TestCreateTaskFromJobFile(t *testing.T) {
	fmt.Println("starting Test for CreateTaskFromJobFile")
	setup()
	fmt.Printf("reading file %s...\n", absPath)
	opts := createATaskFromJobFile(absPath)
	if !(opts.Debug && opts.Load && opts.Help) || opts.Version {
		fmt.Fprintf(os.Stderr, "parsed from json: %v", opts)
		t.Fail()
	}
	tearDown()
	fmt.Println("ending Test for CreateTaskFromJobFile")
}

func tearDown() {
	absPath := path.Join(os.TempDir(), toUnMarshall)
	err := os.Remove(absPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not remove the file %s after a test", absPath)
	}
}
