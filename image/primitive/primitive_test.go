package primitive

import (
	"io"
	"os"
	"testing"

	"github.com/CloudBroker/dash_utils/dashtest"
)

func TestTranform(t *testing.T) {
	var image io.Reader
	image, _ = os.Open("input.png")
	image, err := Transform(image, ".png", "2", "100")
	if err != nil {
		t.Error("Expected reader got", err)
	}
}
func TestPrimitive(t *testing.T) {
	_, err := primitive("-i in.png -o out.png -n 20 -m 0", "out.png")
	if err == nil {
		t.Error("Expected img got", err)
	}
	_, err = primitive("-i input.png -o output.png -n 10 -m 0", "out.png")
	if err == nil {
		t.Error("Expected img got", err)
	}
}

func TestTempFile(t *testing.T) {
	_, err := tempfile("", "")
	if err != nil {
		t.Error("error in creating tempFile")
	}
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
