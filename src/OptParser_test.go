package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func setupOptParserTest() (string, error) {
	toUnMarshall := "job.js"
	// "-s=\"showver.sh\"", "-t=\"useless_template\"", "-d", "-p", "app"
	json := `{
  "debug"      : true,
  "help"       : false,
  "load"       : false,
  "version"    : false,
  "command"    : "",
  "pretty"     : true,
  "scriptName" : "\"showver.sh\"",
  "template"   : "\"useless_template\"",
  "planets"    : [ "app"]
}
`
	data := []byte(json)
	absPath := path.Join(os.TempDir(), toUnMarshall)

	if err := ioutil.WriteFile(absPath, data, 0644); err != nil {
		return "", err
	}
	return absPath, nil
}

func TestCreateTaskFromJobFile(t *testing.T) {
	fmt.Println("starting Test for CreateTaskFromJobFile")
	var absPath string
	var err error
	if absPath, err = setupOptParserTest(); err != nil {
		t.Fail()
	}

	fmt.Printf("reading file %s...\n", absPath)
	opts := createATaskFromJobFile(absPath)
	if opts.Help || opts.Version {
		fmt.Fprintf(os.Stderr, "parsed from json: %v", opts)
		t.Fail()
	}
	// tearDownOptParserTest(absPath)
	fmt.Println("ending Test for CreateTaskFromJobFile")
}

func tearDownOptParserTest(absPath string) {
	if err := os.Remove(absPath); err != nil {
		fmt.Fprintf(os.Stderr, "Could not remove the file %s after a test", absPath)
	}
}
