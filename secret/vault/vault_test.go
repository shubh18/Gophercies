package secret

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/CloudBroker/dash_utils/dashtest"
	homedir "github.com/mitchellh/go-homedir"
)

func secretpath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, "test.txt")
}

func TestSet(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	err := vault.SetKey("hello", "testing")
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}
func TestSetNegative(t *testing.T) {
	file := secretpath()
	vault := NewVault("", file)
	err := vault.SetKey("demo", "testing")
	if err == nil {
		t.Error("Expected  Error but got nil")
	}
}

func TestGet(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	_, err := vault.GetValue("hello")
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}
func TestGetNegative(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	_, err := vault.GetValue("abc")
	if err == nil {
		t.Error("Expected Error but got nil")
	}
	vault = NewVault("", file)
	_, err = vault.GetValue("abc")
	if err == nil {
		t.Error("Expected Error but got nil ")
	}
}

func TestLoad(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	err := vault.load()
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func TestLoadNegative(t *testing.T) {
	home, _ := homedir.Dir()
	file := filepath.Join(home, "")
	vault := NewVault("abc", file)
	err := vault.load()
	if err == nil {
		t.Error("Expected error but got nil", err)
	}
	os.Remove(file)
}
func TestSave(t *testing.T) {
	var v Vault
	err := v.save()
	if err == nil {
		t.Error("Expected error but got nil ")
	}
	deleteFile()
}

func deleteFile() {
	file := secretpath()
	os.Remove(file)
}
func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
