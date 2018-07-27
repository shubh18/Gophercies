package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"task/db"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

/*
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
*/

func TestListCommand(t *testing.T) {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "cmd.db")
	dbconnect, _ := db.InitDB(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{""}
	listCommand.Run(listCommand, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "List of Tasks:")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	dbconnect.Close()

}
