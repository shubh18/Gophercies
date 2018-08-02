package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"task/db"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestDoCommand(t *testing.T) {
	dbconnect := startDB()
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
	file.Close()
	dbconnect.Close()

}

func TestDoNegative(t *testing.T) {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "my.db")
	dbconnect, _ := db.InitDB(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	testCase := []struct {
		arg      []string
		expected string
		msg      string
	}{
		{
			[]string{"1"},
			"Invalid task number: 1",
			"they must be equal",
		},
		{
			[]string{"a"},
			"",
			"they should be equal",
		},
	}

	for _, test := range testCase {
		doCommand.Run(doCommand, test.arg)
		file.Seek(0, 0)
		content, err := ioutil.ReadAll(file)
		if err != nil {
			t.Error("error occured while test case : ", err)
		}
		val, _ := regexp.Match(test.expected, content)
		assert.Equalf(t, true, val, "they should be equal")
		file.Truncate(0)
		file.Seek(0, 0)
	}
	os.Stdout = oldStdout
	file.Close()
	dbconnect.Close()
}
func TestDoCmdNegativeDB(t *testing.T) {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "tasks.db")
	dbConncet, _ := db.InitDB(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"1"}
	dbConncet.Close()
	doCommand.Run(doCommand, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "error occured")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()

}
