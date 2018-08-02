package main

import (
	"os"
	"testing"

	"github.com/CloudBroker/dash_utils/dashtest"
	homedir "github.com/mitchellh/go-homedir"
)

func TestM(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("following error occured while the main function was executed : ")
		}
	}()
	main()
}

func clearFile() {
	dir, _ := homedir.Dir()
	dir = dir + "/task_db.db"
	file, _ := os.OpenFile(dir, os.O_TRUNC, 0666)
	file.Truncate(0)
	file.Close()
}

func TestMain(m *testing.M) {
	clearFile()
	dashtest.ControlCoverage(m)
	m.Run()
	clearFile()
}
