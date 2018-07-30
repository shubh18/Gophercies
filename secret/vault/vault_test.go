package secret

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

//TestFile basic test case
func TestVault(t *testing.T) {
	home, _ := homedir.Dir()
	filePath := filepath.Join(home, ".test.secrets")

	v := NewVault("test", filePath)
	assert.Equal(t, "test", v.encodingKey)

	key1 := "key1"
	value1 := "value1"
	_ = v.SetKey(key1, value1)

	answer1, _ := v.GetValue(key1)
	assert.Equal(t, answer1, value1)

	s, err := v.GetValue("")
	if err == nil {
		t.Error("Expected nil got", s)
	}

}
