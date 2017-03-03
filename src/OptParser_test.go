package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func setupOptParserTest() string {
	toUnMarshall := "job.js"
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
	data := []byte(json)
	absPath := path.Join(os.TempDir(), toUnMarshall)
	err := ioutil.WriteFile(absPath, data, 0644)
	if err != nil {
		panic(err)
	}
	return absPath
}

func TestCreateTaskFromJobFile(t *testing.T) {
	fmt.Println("starting Test for CreateTaskFromJobFile")
	absPath := setupOptParserTest()
	fmt.Printf("reading file %s...\n", absPath)
	opts := createATaskFromJobFile(absPath)
	if !(opts.Debug && opts.Load && opts.Help) || opts.Version {
		fmt.Fprintf(os.Stderr, "parsed from json: %v", opts)
		t.Fail()
	}
	tearDownOptParserTest(absPath)
	fmt.Println("ending Test for CreateTaskFromJobFile")
}

func tearDownOptParserTest(absPath string) {
	err := os.Remove(absPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not remove the file %s after a test", absPath)
	}
}
