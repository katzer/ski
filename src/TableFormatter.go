package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// TFAdapter ...
type TFAdapter struct {
	real *TableFormatter
}

func (tfAdapter TFAdapter) init() {
	tfAdapter.real.init()
}

func (tfAdapter TFAdapter) format(planets []Planet, opts *Opts, writer io.Writer) {
	tfAdapter.real.format(planets, opts, writer)
}

// TableFormatter prints input in tabular format
type TableFormatter struct {
}

func (tableFormatter *TableFormatter) init() {

}

func (tableFormatter *TableFormatter) format(planets []Planet, opts *Opts, writer io.Writer) {
	for _, planet := range planets {
		formatted := tableFormatter.formatPlanet(planet, opts)
		fmt.Fprint(writer, formatted)
	}
}

func (tableFormatter *TableFormatter) formatPlanet(planet Planet, opts *Opts) string {
	var err error
	err = tableFormatter.writeTmpTable(planet, planet.outputStruct.output)
	if err != nil {
		return err.Error()
	}
	var textFSMOutput string
	textFSMOutput, err = tableFormatter.executeTextFSM(planet, opts)
	if err != nil {
		return err.Error()
	}
	jsonString := convertToJSON(textFSMOutput)
	err = tableFormatter.deleteTmpTable(planet)
	if err != nil {
		return err.Error()
	}
	return strings.Replace(jsonString, "'", "\"", -1)
}

// Converts the quite special format textFSM returns to proper JSON format
func convertToJSON(toConvert string) string {
	return fmt.Sprintf("[\n%s]\n", strings.Replace(toConvert, "]\n[", "],\n[", -1))
}

// executes phyton2 program "textFSM" with provided template and temporary file and returns the answer
func (tableFormatter *TableFormatter) executeTextFSM(planet Planet, opts *Opts) (string, error) {
	tmpTableFile := path.Join(os.Getenv("ORBIT_HOME"), "tmp", tmpTableFileName)
	tmpTableFile = fmt.Sprintf("%s%d.txt", tmpTableFile, planet.outputStruct.position)
	templateFile := path.Join(os.Getenv("ORBIT_HOME"), templateDirectory, opts.Template)
	pyScriptFile := path.Join(os.Getenv("ORBIT_HOME"), thirdPartySoftwareDirectory, textFSMDirectory, textFSMName)

	cmd := exec.Command("python2", pyScriptFile, templateFile, tmpTableFile)
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		message := "thrown from tableFormatter.format->exec pythonscript"
		full := fmt.Sprintf("%s\n --- Additional info: %s\n", err, message)
		planet.outputStruct.output = fmt.Sprintf("%s\n%s", planet.outputStruct.output, full)
		planet.outputStruct.errored = true
		log.Warn(full)
		return "", err
	}
	formattedString := strings.Split(out.String(), "FSM Table:\n")[1]
	return formattedString, nil
}

// Writes the provided string to a temporary file
func (tableFormatter *TableFormatter) writeTmpTable(planet Planet, toWrite string) error {

	tempdir := path.Join(os.Getenv("ORBIT_HOME"), "tmp")
	tmpTableFile := path.Join(tempdir, tmpTableFileName)
	tmpTableFile = fmt.Sprintf("%s%d.txt", tmpTableFile, planet.outputStruct.position)
	err := ioutil.WriteFile(tmpTableFile, []byte(toWrite), 0644)
	if err != nil {
		message := fmt.Sprintf("Attempt to write a temporary file for textfsm execution failed: %s\n", tmpTableFile)
		full := fmt.Sprintf("%s\n --- Additional info: %s\n", err, message)
		planet.outputStruct.output = fmt.Sprintf("%s\n%s", planet.outputStruct.output, full)
		planet.outputStruct.errored = true
		log.Errorln(full)
		return err
	}
	return nil
}

// deletes temporary file needed for textFSM
func (tableFormatter *TableFormatter) deleteTmpTable(planet Planet) error {
	tmpTableFile := path.Join(os.Getenv("ORBIT_HOME"), "tmp", tmpTableFileName)
	tmpTableFile = fmt.Sprintf("%s%d.txt", tmpTableFile, planet.outputStruct.position)
	err := os.Remove(tmpTableFile)
	if err != nil {
		message := fmt.Sprintf("Attempt to delete the temporary file for textfsm execution failed: %s\n", tmpTableFile)
		full := fmt.Sprintf("%s\n --- Additional info: %s\n", err, message)
		planet.outputStruct.output = fmt.Sprintf("%s\n%s", planet.outputStruct.output, full)
		planet.outputStruct.errored = true
		log.Errorln(full)
		return err
	}
	return nil
}
