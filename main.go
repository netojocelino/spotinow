package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/netojocelino/spotinow/cli"
)

func openURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		// cmd = "xdg-open"
		cmd = "firefox"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Welcome")
		fmt.Println("Spotify CLI - Saída o que está ouvindo.")
		return
	}

	command := os.Args[1]

	cli.HandlerCommandLineInput(command)

}
