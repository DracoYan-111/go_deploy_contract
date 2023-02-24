package aescrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
)

// Encrypt encrypts a string with a given key.
func Encrypt(plaintext string, key []byte) (string, error) {

	// 填充数据
	paddingLength := aes.BlockSize - len(plaintext)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	plaintext = plaintext + string(padding)

	// 创建一个新的AES加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 加密数据
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:i+aes.BlockSize], []byte(plaintext[i:i+aes.BlockSize]))
	}

	// 将加密后的数据转换为base64字符串
	result := strings.TrimRight(base64.StdEncoding.EncodeToString(ciphertext), "=")

	return result, err
}

// AesDecrypt Decrypt decrypts a string with a given key.
func AesDecrypt(cipher string, key []byte) (string, error) {
	// 密钥
	//key := []byte("0123456789abcdef")

	// 需要解密的密文
	ciphertext, _ := base64.StdEncoding.DecodeString(cipher)

	// 创建一个新的AES解密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建一个ECB模式的解密器
	mode := newECBDecrypter(block)

	// 解密数据
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去掉填充的部分
	plaintext = pKCS7Unpad(plaintext, block.BlockSize())

	return string(plaintext), err
}

// NewECBDecrypter 创建一个新的ECB模式解密器
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

// PKCS7Unpad 去除PKCS7填充
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
