package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"strings"
)

// Encrypt Encrypts a string with a given key.
func Encrypt(plaintext string) (string, error) {
	loading, err := ConfigurationLoading("server", []string{"key"})
	if err != nil {
		log.Println(err)
	}
	key := []byte(loading[0])

	paddingLength := aes.BlockSize - len(plaintext)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	plaintext = plaintext + string(padding)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:i+aes.BlockSize], []byte(plaintext[i:i+aes.BlockSize]))
	}

	result := strings.TrimRight(base64.StdEncoding.EncodeToString(ciphertext), "=")

	return result, err
}

// AesDecrypt Decrypt decrypts a string with a given key.
func AesDecrypt(cipher string) (string, error) {
	loading, err := ConfigurationLoading("server", []string{"key"})
	if err != nil {
		log.Println(err)
	}
	key := []byte(loading[0])

	// The ciphertext to be decrypted
	ciphertext, _ := base64.StdEncoding.DecodeString(cipher)

	// Create a new AES decryption block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a decryptor in ECB mode
	mode := newECBDecrypter(block)

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove the filled part
	plaintext = pKCS7Unpad(plaintext, block.BlockSize())

	return string(plaintext), err
}

// NewECBDecrypter Create a new ECB mode decryptor
func newECBDecrypter(block cipher.Block) cipher.BlockMode {
	return &ecbDecrypter{block}
}

type ecbDecrypter struct {
	b cipher.Block
}

func (x *ecbDecrypter) BlockSize() int { return x.b.BlockSize() }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.b.BlockSize() != 0 {
		panic("input not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst[:x.b.BlockSize()], src[:x.b.BlockSize()])
		src = src[x.b.BlockSize():]
		dst = dst[x.b.BlockSize():]
	}
}

// PKCS7Unpad Remove PKCS 7 padding
func pKCS7Unpad(data []byte, blockSize int) []byte {
	if len(data) == 0 {
		return []byte{}
	}
	padding := int(data[len(data)-1])
	if padding < 1 || padding > blockSize {
		return data
	}
	for i := len(data) - 1; i > len(data)-padding-1; i-- {
		if int(data[i]) != padding {
			return data
		}
	}
	return data[:len(data)-padding]
}
