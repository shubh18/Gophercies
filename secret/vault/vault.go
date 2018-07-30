package secret

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"secret/cipher"
)

// Vault is public struct
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// NewVault is thin simple method for creating new Vault
func NewVault(key, path string) *Vault {
	return &Vault{
		encodingKey: key,
		filepath:    path,
	}
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

//It loads vault and if the map does not exist it creates one using calling make()
func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()

	reader, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}

	return v.readKeyValues(reader)
}

//save function writes encoding key to vault
func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

//SetKey is used for storing key:value pair in Vault.
//It is thread safe.
func (v *Vault) SetKey(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	err = v.save()
	return err
}

//GetValue is used for retriving value for a specified key.
func (v *Vault) GetValue(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	return value, nil
}
