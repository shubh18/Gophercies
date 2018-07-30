package cipher

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"

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
func TestEncrypt(t *testing.T) {
	cipher, err := Encrypt("Tendulkar", "Sachin tendulkar")
	if err != nil {
		t.Error("Expected nil got", err)

	}
	fmt.Println(cipher)
}

func TestDecrypt(t *testing.T) {
	cipher, err := Encrypt("Tendulkar", "Sachin tendulkar")
	plain, err := Decrypt("Tendulkar", cipher)
	if err != nil {
		t.Error("Expected nil got", err)

	}
	fmt.Println(plain)
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
	reader, err := DecryptReader("test", f)
	if err != nil {
		t.Error("Expected reader got", reader)
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
	writer, err := EncryptWriter(v.encodingKey, f)
	if err != nil {
		t.Error("Expected Writer got", writer)
	}

}
