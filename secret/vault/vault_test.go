package secret

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func secretpath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".test.secrets")
}
func TestSet(t *testing.T) {
	testSuit := []struct {
		encodingKey string
		filepath    string
		key         string
		plainText   string
	}{
		{encodingKey: "demokey", filepath: secretpath(), key: "secret", plainText: "hello"},
	}
	for _, test := range testSuit {
		v := NewVault(test.encodingKey, test.filepath)
		err := v.SetKey(test.key, test.plainText)
		if err != nil {
			t.Error("error in Set")
		}
	}
}

func TestGet(t *testing.T) {
	testSuit := []struct {
		encodingKey string
		filepath    string
		key         string
		plainText   string
	}{
		{encodingKey: "demokey", filepath: secretpath(), key: "secret", plainText: "hello"},
		{encodingKey: "demokey", filepath: secretpath(), key: "picasso", plainText: ""},
	}
	for _, test := range testSuit {
		v := NewVault(test.encodingKey, test.filepath)
		plainText, _ := v.GetValue(test.key)
		if plainText != test.plainText {
			t.Error("error in Get")
		}
	}
}
