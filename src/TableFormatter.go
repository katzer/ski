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

const pythonScriptName = "textfsm.py"

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

	pyScriptFile := path.Join(opts.pyScriptPath, pythonScriptName)
	/**
	if runtime.GOOS == "windows" {
		pyScriptFile = os.Getenv("TEMP") + "\\tempTabFormat.py"
	} else {
		pyScriptFile = os.Getenv("HOME") + "/tempTabFormat.py"
	}*/
	cmd := exec.Command("python2", pyScriptFile, templateFile, tmpTableFile)
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Printf("toFormat: %s\n", toFormat)
		throwErrExt(err, "thrown from tableFormatter.format->exec pythonscript")
	}
	formattedString := strings.Split(out.String(), "FSM Table:\n")[1]
	formattedString = strings.TrimSpace(formattedString)
	cleanString := tableFormatter.cleanEntries(formattedString)

	err = os.Remove(tmpTableFile)
	if err != nil {
		log.Fatal(err)
	}

	//################################# temporaray hardocode##############################################
	cleanString = cleanifyTable(cleanString)
	cleanString = fmt.Sprintf("[\n%s]\n", strings.Replace(cleanString, "]\n[", "],\n[", -1))
	//################################# temporaray hardocode##############################################

	return cleanString

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

func (tableFormatter *TableFormatter) cleanEntries(toParse string) string {
	split := strings.Split(toParse, "\n")
	split = split[:len(split)-1]
	cleaned := ""
	for _, entry := range split {
		row := strings.Split(entry, ", ")
		row[0] = strings.TrimPrefix(row[0], "[")
		row[0] = strings.TrimPrefix(row[0], "'")
		row[0] = strings.TrimSuffix(row[0], "'")
		row[0] = strings.TrimSpace(row[0])
		row[0] = fmt.Sprintf("['%s'", row[0])
		row[1] = strings.TrimSuffix(row[1], "]")
		row[1] = strings.TrimPrefix(row[1], "'")
		row[1] = strings.TrimSuffix(row[1], "'")
		row[1] = strings.TrimSpace(row[1])
		row[1] = fmt.Sprintf("'%s']", row[1])
		cleaned = fmt.Sprintf("%s%s, %s\n", cleaned, row[0], row[1])
	}

	return cleaned
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
