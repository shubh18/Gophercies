package commands

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestAddCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "add", "Watch Blacklist")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

}

func TestAddCommandHelp(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "add", "-h")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

}
