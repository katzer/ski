package main

import (
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

// TableFormatter prints input in tabular format
type TableFormatter struct {
}

func (tableFormatter *TableFormatter) format(planet Planet, opts *Opts) string {
	tableFormatter.writeTmpTable(planet.outputStruct.output)
	jsonString := convertToJSON(tableFormatter.executeTextFSM(planet, opts))
	tableFormatter.deleteTmpTable()
	return strings.Replace(jsonString, "'", "\"", -1)

}

// Converts the quite special format textFSM returns to proper JSON format
func convertToJSON(toConvert string) string {
	return fmt.Sprintf("[\n%s]\n", strings.Replace(toConvert, "]\n[", "],\n[", -1))
}

// executes phyton2 program "textFSM" with provided template and temporary file and returns the answer
func (tableFormatter *TableFormatter) executeTextFSM(planet Planet, opts *Opts) string {
	tmpTableFile := path.Join(os.Getenv("ORBIT_HOME"), tmpTableFileName)
	templateFile := path.Join(os.Getenv("ORBIT_HOME"), templateDirectory, opts.template)
	pyScriptFile := path.Join(os.Getenv("ORBIT_HOME"), thirdPartySoftwareDirectory, textFSMDirectory, textFSMName)

	cmd := exec.Command("python2", pyScriptFile, templateFile, tmpTableFile)
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		message := "thrown from tableFormatter.format->exec pythonscript"
		full := fmt.Sprintf("%s\n --- Additional info: %s\n", err, message)
		os.Stderr.WriteString(full)
		log.Errorln(full)
		log.Fatalf("Format: %s\n", planet.outputStruct.output)
	}
	formattedString := strings.Split(out.String(), "FSM Table:\n")[1]
	return formattedString
}

// Writes the provided string to a temporary file
func (tableFormatter *TableFormatter) writeTmpTable(toWrite string) {
	tmpTableFile := path.Join(os.Getenv("ORBIT_HOME"), tmpTableFileName)
	err := ioutil.WriteFile(tmpTableFile, []byte(toWrite), 0644)
	if err != nil {
		log.Fatalf("Attempt to write a temporary file for textfsm execution failed: %s\n", tmpTableFile)
	}

}

// deletes temporary file needed for textFSM
func (tableFormatter *TableFormatter) deleteTmpTable() {
	tmpTableFile := path.Join(os.Getenv("ORBIT_HOME"), tmpTableFileName)
	err := os.Remove(tmpTableFile)
	if err != nil {
		log.Fatal(err)
	}

}
