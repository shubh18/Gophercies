package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"task/db"
	"testing"

	"github.com/CloudBroker/dash_utils/dashtest"
	"github.com/boltdb/bolt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func startDB() *bolt.DB {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "cmd.db")
	dbconnect, _ := db.InitDB(DbPath)
	return dbconnect
}

func TestAddCommand(t *testing.T) {
	dbconnect := startDB()
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"Watch Golang tutorial"}
	addCommand.Run(addCommand, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Added")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()
	dbconnect.Close()

}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
