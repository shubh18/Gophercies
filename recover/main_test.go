package main

import (
	"net/http"
	"os"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"

	"github.com/stretchr/testify/assert"
)

func TestM(t *testing.T) {
	templistenAndServe := listenAndServeFunc
	defer func() {
		listenAndServeFunc = templistenAndServe
	}()
	listenAndServeFunc = func(port string, hanle http.Handler) error {
		panic("testing")
	}
	assert.PanicsWithValuef(t, "testing", main, "they should be equal")
}

func TestMain(m *testing.M) {
	file, _ := os.OpenFile("testing.txt", os.O_CREATE, 0000)
	file.Close()
	dashtest.ControlCoverage(m)
	m.Run()
	os.Remove("testing.txt")
}
