package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

var version = "undefined"

func printVersion() {
	runtimeOS := getOS()
	progArch := getArch()
	archOS := getOSArch()
	vers := fmt.Sprintf("ski version %s %s %s (%s)", version, progArch, runtimeOS, archOS)
	fmt.Printf("%s\n", vers)
}

func getOS() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "linux":
		return "Linux"
	case "darwin":
		return "MacOS"
	default:
		return "could not determine OS"
	}
}

func getArch() string {
	switch runtime.GOARCH {
	case "amd64":
		return "64bit"
	case "386":
		return "32bit"
	default:
		return "could not determine architecture"
	}
}

func getOSArch() string {
	switch runtime.GOOS {
	case "linux":
		out, err := exec.Command("uname", "-m").Output()
		if err != nil {
			fmt.Println("error occured")
			fmt.Printf("%s", err)
		}
		return strings.TrimSuffix(string(out), "\n")
	case "windows":
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
