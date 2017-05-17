package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

var version = "undefined"

func printVersion() {
	goos := getGOOS()
	goarch := getGOARCH()
	osArch := getOSArch()
	vers := fmt.Sprintf("v%s - %s %s (%s)", version, goos, goarch, osArch)
	fmt.Printf("%s\n", vers)
}

func getGOOS() string {
	switch runtime.GOOS {
	case windows:
		return "Windows_NT"
	case linux:
		return "Linux"
	case mac:
		return "Darwin"
	default:
		return "could not determine OS"
	}
}

func getGOARCH() string {
	switch runtime.GOARCH {
	case "amd64":
		return "64-Bit"
	case "386":
		return "32-Bit"
	case "686":
		return "32-Bit"
	default:
		return "could not determine architecture"
	}
}

func getOSArch() string {
	switch runtime.GOOS {
	case linux:
		out, err := exec.Command("uname", "-m").Output()
		if err != nil {
			fmt.Println("error occured")
			fmt.Printf("%s", err)
		}
		return strings.TrimSuffix(string(out), "\n")
	case windows:
		out, err := exec.Command("echo %PROCESSOR_ARCHITECTURE%").Output()
		if err != nil {
			return "could not determine OS-architecture"
		}
		return strings.TrimSuffix(string(out), "\n")
	default:
		return "could not determine Operating system"
	}
}
