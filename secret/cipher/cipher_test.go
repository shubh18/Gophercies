package cipher

import (
	"fmt"
	"testing"
)

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
