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

/*func TestDoCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "do", "1")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

}
*/
func TestDoCommand(t *testing.T) {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "cmd.db")
	dbconnect, _ := db.InitDB(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"1"}
	doCommand.Run(doCommand, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Marked")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	dbconnect.Close()

}
