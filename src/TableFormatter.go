package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// TableFormatter prints input in tabular format
type TableFormatter struct {
}

func (tableFormatter *TableFormatter) format(toFormat string, opts *Opts) string {
	tableFormatter.writeTmpTable(toFormat)
	jsonString := convertToJSON(tableFormatter.executeTextFSM(toFormat, opts))
	tableFormatter.deleteTmpTable()
	return strings.Replace(jsonString, "'", "\"", -1)

}

func (tableFormatter *TableFormatter) cleanEntries(toParse string) string {
	split := strings.Split(toParse, "\n")
	split = split[:len(split)-1]
	cleaned := ""
	for _, entry := range split {
		row := strings.Split(entry, ", ")
		row[0] = tableFormatter.cleanEntry(row[0], true)
		row[1] = tableFormatter.cleanEntry(row[1], false)
		cleaned = fmt.Sprintf("%s%s, %s\n", cleaned, row[0], row[1])
	}

	return cleaned
}

func (tableFormatter *TableFormatter) cleanEntry(row string, key bool) string {
	cleanedComponent := ""
	if key {
		cleanedComponent = strings.TrimPrefix(row, "[")
	} else {
		cleanedComponent = strings.TrimPrefix(row, "]")
	}
	cleanedComponent = strings.TrimPrefix(cleanedComponent, "'")
	cleanedComponent = strings.TrimSuffix(cleanedComponent, "'")
	cleanedComponent = strings.TrimSpace(cleanedComponent)
	if key {
		cleanedComponent = fmt.Sprintf("['%s'", cleanedComponent)
	} else {
		cleanedComponent = fmt.Sprintf("'%s']", cleanedComponent)
	}
	return cleanedComponent
}

func cleanifyTable(toclean string) string {
	split := strings.Split(toclean, "\n")
	split = split[:len(split)-1]
	cleaned := ""
	FAST := false
	for _, entry := range split {
		row := strings.Split(entry, ", ")
		if FAST {
			row[0] = fmt.Sprintf("['%s'", strings.TrimSpace(strings.Split(row[1], "|")[2]))
			row[1] = fmt.Sprintf("'%s']", strings.TrimSpace(strings.Split(row[1], "|")[3]))
		}
		if strings.Contains(row[0], "-[FAST") {
			FAST = true
			continue
		}
		cleaned = fmt.Sprintf("%s%s, %s\n", cleaned, row[0], row[1])
	}
	return cleaned
}

func convertToJSON(toConvert string) string {
	return fmt.Sprintf("[\n%s]\n", strings.Replace(toConvert, "]\n[", "],\n[", -1))
}

func (tableFormatter *TableFormatter) executeTextFSM(toFormat string, opts *Opts) string {
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
		log.Fatalf("Format: %s\n", toFormat)
	}
	formattedString := strings.Split(out.String(), "FSM Table:\n")[1]
	return formattedString
}

func (tableFormatter *TableFormatter) writeTmpTable(toWrite string) {
	tmpTableFile := path.Join(os.Getenv("ORBIT_HOME"), tmpTableFileName)
	err := ioutil.WriteFile(tmpTableFile, []byte(toWrite), 0644)
	if err != nil {
		log.Fatalf("Attempt to write a temporary file for textfsm execution failed: %s\n", tmpTableFile)
	}

}

func (tableFormatter *TableFormatter) deleteTmpTable() {
	tmpTableFile := path.Join(os.Getenv("ORBIT_HOME"), tmpTableFileName)
	err := os.Remove(tmpTableFile)
	if err != nil {
		log.Fatal(err)
	}

}
