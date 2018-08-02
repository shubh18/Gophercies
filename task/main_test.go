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

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
	m.Run()
	
}
