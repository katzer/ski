package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func parseConnectionDetails(planetID string) Planet {
	skiString := getFullSkiString(planetID)
	var planet Planet
	tokens := strings.Split(skiString, skiDelim)
	planet.planetType = tokens[0]
	connectionURL := tokens[len(tokens)-1]
	urlTokens := strings.Split(connectionURL, ":")
	if len(urlTokens) > 1 {
		planet.dbID = urlTokens[0]
	}
	planet.user, planet.host = getUserAndHost(connectionURL)
	log.Debugf("skiString: %s, and planet parsed from it: %v", planet)
	return planet
}

func getKeyPath() string {
	keyPath := os.Getenv("ORBIT_KEY")
	if keyPath == "" {
		if runtime.GOOS == "windows" {
			keyPath = os.Getenv("TEMP") + "\\tempTabFormat.py"
		} else {
			keyPath = path.Join(os.Getenv("ORBIT_HOME"), "config", "ssh", "orbit.key")
		}
	}
	return strings.TrimPrefix(keyPath, os.Getenv("HOME"))
}

func isSupported(planetType string) bool {
	supported := map[string]bool{database: true, linuxServer: true, webServer: false}
	return supported[planetType]
}

func getType(skiString string) string {
	return strings.Split(skiString, skiDelim)[0]
}

func getFullSkiString(id string) string {
	cmd := exec.Command("fifa", "-f=ski", id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		message := fmt.Sprintf("%s output is: %s called from ErrOut.\n", err, out)
		os.Stderr.WriteString("Unknown target\n")
		log.Fatalf(message)
	}
	// TODO fifa sends a newline
	return strings.TrimSuffix(string(out), "\n")
}

func getUserAndHost(connectionURL string) (string, string) {
	var tokens []string
	idx := strings.IndexRune(connectionURL, ':')
	if idx < 0 {
		tokens = strings.Split(connectionURL, "@")
		return tokens[0], tokens[1]
	}
	tokens = strings.Split(connectionURL[idx+1:], "@")
	return tokens[0], tokens[1]
}
