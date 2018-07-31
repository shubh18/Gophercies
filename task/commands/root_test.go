package commands

import (
	"testing"
)

func TestRoot(t *testing.T) {
	err := RootCmd.Execute()
	if err != nil {
		t.Error("error in root command")
	}
}
