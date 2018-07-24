package main

import (
	"testing"
)

func TestInitApplication(t *testing.T) {
	err := initApplication()
	if err != nil {
		t.Error("Expected nil got,", err)
	}

}
