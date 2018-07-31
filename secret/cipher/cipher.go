package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	var cipherKey cipher.Stream
	block, err := newCipherBlock(key)
	if err == nil {
		cipherKey = cipher.NewCFBEncrypter(block, iv)
	}
	return cipherKey, err
}

// EncryptWriter will return a writer that will write encrypted data to
// the original writer.
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)
	stream, _ := encryptStream(key, iv)
	n, err := w.Write(iv)
	if n != len(iv) || err != nil {
		return nil, errors.New("encrypt: unable to write full iv to writer")
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	var cipherKey cipher.Stream
	block, err := newCipherBlock(key)
	if err == nil {
		cipherKey = cipher.NewCFBDecrypter(block, iv)
	}
	return cipherKey, err
}

// DecryptReader will return a reader that will decrypt data from the
// provided reader and give the user a way to read that data as it if was
// not encrypted.
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	//apart from nil check we should also ensure that
	//number of bytes that are read must be 16
	if n < len(iv) || err != nil {
		return nil, errors.New("encrypt: unable to read the full iv")
	}
	stream, err := decryptStream(key, iv)
	return &cipher.StreamReader{S: stream, R: r}, err
}

//newCipherBlock return cipher block containing hashed version of key
func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
