package commands

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestListCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "list")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

}

func TestListCommandHelp(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "list", "-h")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

}
