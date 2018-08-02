package hyperlink

import (
	"runtime/debug"
	"testing"

	"github.com/CloudBroker/dash_utils/dashtest"
)

func TestCreateLinks(t *testing.T) {
	stack := debug.Stack()
	link := CreateLinks(string(stack))
	if link == "" {
		t.Error("Expected link got", link)
	}
}
func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
