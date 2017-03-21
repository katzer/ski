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
	vers := fmt.Sprintf("ski version %s (%s %s)", version, goos, goarch)
	fmt.Printf("%s\n", vers)
}

func getGOOS() string {
	switch runtime.GOOS {
	case windows:
		return "Windows"
	case linux:
		return "Linux"
	case mac:
		return "MacOS"
	default:
		return "could not determine OS"
	}
}

func getGOARCH() string {
	switch runtime.GOARCH {
	case "amd64":
		return "64bit"
	case "386":
		return "32bit"
	case "686":
		return "32bit"
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
		out, err := exec.Command("if exist \"%ProgramFiles(x86)%\" echo 64-bit").Output()
		if err != nil {
			fmt.Println("error occured")
			fmt.Printf("%s", err)
		}
		if string(out) == "64-bit" {
			return "x86_64"
		}
		return "i686"
	default:
		return "could not determine Operating system"
	}
}
