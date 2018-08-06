package cipher

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/CloudBroker/dash_utils/dashtest"
	homedir "github.com/mitchellh/go-homedir"
)

type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

func NewVault(key, path string) *Vault {
	return &Vault{
		encodingKey: key,
		filepath:    path,
	}
}

func TestDecryptReader(t *testing.T) {
	home, _ := homedir.Dir()
	file := filepath.Join(home, ".test.secrets")
	v := NewVault("test", file)
	f, err := os.Open(file)
	if err != nil {
		v.keyValues = make(map[string]string)
	}
	defer f.Close()
	_, err = DecryptReader("test", f)
	if err != nil {
		t.Error("Expected reader got", err)
	}

}

func TestEncryptWriter(t *testing.T) {
	home, _ := homedir.Dir()
	file := filepath.Join(home, ".test.secrets")
	v := NewVault("test", file)
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("TestEncryptWriter:", err)
	}
	defer f.Close()
	_, err = EncryptWriter(v.encodingKey, f)
	if err != nil {
		t.Error("Expected Writer got", err)
	}

}

func TestDecryptReaderNegative(t *testing.T) {
	home, _ := homedir.Dir()
	file := filepath.Join(home, "test_secret.txt")

	f, _ := os.Open(file)
	defer f.Close()
	_, err := DecryptReader("abc", f)
	if err == nil {
		t.Error("Expected error but got no error")
	}
	os.Remove(file)
}

func TestEncryptWriterNegative(t *testing.T) {
	home, _ := homedir.Dir()
	file := filepath.Join(home, ".test.secrets")
	v := NewVault("test", file)
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("TestEncryptWriter:", err)
	}
	defer f.Close()
	writer, err := EncryptWriter("", f)
	if err != nil {
		t.Error("Expected Writer got", writer)
	}

}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
