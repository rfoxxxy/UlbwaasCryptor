package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

func encryptStringToHEX(enc string) string {
	hx := hex.EncodeToString([]byte(enc))
	return fmt.Sprintf("%v", hx)
}

func encrypt(key []byte, message []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	b := message
	b = PKCS5Padding(b, aes.BlockSize)
	encMessage := make([]byte, len(b))
	iv := key[:aes.BlockSize]
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encMessage, b)

	return encMessage, nil
}

func decrypt(key []byte, encMessage []byte) ([]byte, error) {
	iv := key[:aes.BlockSize]
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	if len(encMessage) < aes.BlockSize {
		return nil, errors.New("encMessage слишком короткий")
	}

	decrypted := make([]byte, len(encMessage))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, encMessage)

	return PKCS5UnPadding(decrypted), nil
}

func PKCS5Padding(cipher []byte, blockSize int) []byte {
	padding := blockSize - len(cipher)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipher, padText...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])

	return src[:(length - unPadding)]
}
