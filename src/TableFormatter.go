package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

type TableFormatter struct {
}

func (tableFormatter *TableFormatter) format(toFormat string, opts *Opts) string {
	tmpTableFile := fmt.Sprintf("%s/orbitTable.txt", os.Getenv("HOME"))
	err := ioutil.WriteFile(tmpTableFile, []byte(toFormat), 0644)
	if err != nil {
		fmt.Println("writefile failed!")
		os.Exit(1)
	}
	templateFile := path.Join(opts.templatePath, opts.templateName)

	pys := getPyScript()
	pyScriptFile := ""
	if runtime.GOOS == "windows" {
		pyScriptFile = os.Getenv("TEMP") + "\\tempTabFormat.py"
	} else {
		pyScriptFile = os.Getenv("HOME") + "/tempTabFormat.py"
	}
	err = ioutil.WriteFile(pyScriptFile, []byte(pys), 0644)
	if err != nil {
		fmt.Println("writing pyscript failed")
		log.Fatal(err)
	}
	cmd := exec.Command("python2", pyScriptFile, templateFile, tmpTableFile)
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("executing pyscript failed")
		log.Fatal(err)
	}
	formattedString := strings.Split(out.String(), "FSM Table:\n")[1]
	formattedString = strings.TrimSpace(formattedString)
	parsedTable := tableFormatter.parseFSMOutput(formattedString)
	fmt.Printf("pased Structure: %v\n", parsedTable)
	fmt.Println("raw table")
	fmt.Printf("%s\n", formattedString)

	err = os.Remove(pyScriptFile)
	if err != nil {
		fmt.Println("removing pyscript failed")
		log.Fatal(err)
	}

	err = os.Remove(tmpTableFile)
	if err != nil {
		log.Fatal(err)
	}
	return ""

}

func (tableFormatter *TableFormatter) parseFSMOutput(toParse string) map[string]string {
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
