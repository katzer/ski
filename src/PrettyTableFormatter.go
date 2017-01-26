package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

const prettyPythonScriptName = "texttable.py"

type PrettyTableFormatter struct {
}

func (prettyTableFormatter *PrettyTableFormatter) format(toFormat string, opts *Opts) string {
	tmpTableFile := fmt.Sprintf("%s/orbitTable.txt", os.Getenv("HOME"))
	err := ioutil.WriteFile(tmpTableFile, []byte(toFormat), 0644)
	if err != nil {
		fmt.Println("writefile failed!")
		os.Exit(1)
	}
	templateFile := path.Join(os.Getenv("ORBIT_HOME"), templateDirectory, opts.template)

	pyScriptFile := path.Join(os.Getenv("ORBIT_HOME"), thirdPartySoftwareDirectory, textFSMDirectory, prettyPythonScriptName)

	cmd := exec.Command("python2", pyScriptFile, templateFile, tmpTableFile)
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Printf("toFormat: %s\n", toFormat)
		throwErrExt(err, "thrown from tableFormatter.format->exec pythonscript")
	}
	//formattedString := strings.Split(out.String(), "FSM Table:\n")[1]
	formattedString := strings.TrimSpace(out.String())

	err = os.Remove(tmpTableFile)
	if err != nil {
		log.Fatal(err)
	}
	return formattedString

}

func (prettyTableFormatter *PrettyTableFormatter) parseFSMOutput(toParse string) map[string]string {
	var parsed map[string]string
	parsed = make(map[string]string)
	split := strings.Split(toParse, "\n")
	split = split[0 : len(split)-2]
	for _, entry := range split {
		row := strings.Split(entry, " ")
		parsed[strings.TrimSuffix(row[0], ",")] = strings.Join(row[1:], "")
	}

	return parsed
}
