package vault

import (
	"errors"
	"fmt"
	encrypt "secret/encrypt"
)

//NewVault creates vault to store keys
func NewVault(encodingKey string) Vault {
	return Vault{
		encodingKey: encodingKey,
		keyValues:   make(map[string]string),
	}
}

//Vault struct is representation of Vault
type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

//Get to get decrypted text
func (v *Vault) Get(key string) (string, error) {
	hex, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret:no value for key")
	}
	ret, err := encrypt.Decrypt(hex, v.encodingKey)
	if err != nil {
		fmt.Println("Content not Decrypted")
	}
	fmt.Println(ret)
	return ret, nil
}

//Set to decrypt plain text
func (v *Vault) Set(key, value string) error {
	encryptedValue, err := encrypt.Encrypt(v.encodingKey, value)
	if err != nil {
		fmt.Println("Nothing to encrypt")

	}
	v.keyValues[key] = encryptedValue
	return nil
}
