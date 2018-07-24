package commands

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestDoCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "do", "1")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

}
func TestDoCommandHelp(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "do", "-h")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

}
