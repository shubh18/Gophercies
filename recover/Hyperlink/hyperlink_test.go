package hyperlink

import (
	"runtime/debug"
	"testing"
)

func TestCreateLinks(t *testing.T) {
	stack := debug.Stack()
	link := CreateLinks(string(stack))
	if link == "" {
		t.Error("Expected link got", link)
	}
}
