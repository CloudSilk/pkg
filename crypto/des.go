package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

// 使用PKCS5进行填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去除PKCS5填充
func PKCS5UnPadding(cipherText []byte, blockSize int) ([]byte, error) {
	length := len(cipherText)
	if length == 0 {
		return nil, fmt.Errorf("cipherText is empty")
	}

	unpadding := int(cipherText[length-1])
	if unpadding > blockSize || unpadding > length {
		return nil, fmt.Errorf("invalid padding size")
	}

	paddingStart := length - unpadding
	for i := paddingStart; i < length; i++ {
		if cipherText[i] != byte(unpadding) {
			return nil, fmt.Errorf("invalid padding")
		}
	}

	return cipherText[:paddingStart], nil
}

// DesEncrypt DES加密
func DesEncrypt(plainText, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plainText = PKCS5Padding(plainText, blockSize)
	cipherText := make([]byte, len(plainText))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(cipherText, plainText)
	return cipherText, nil
}

// DesDecrypt DES解密
func DesDecrypt(cipherText, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(cipherText) < blockSize {
		return nil, fmt.Errorf("cipherText too short")
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	return PKCS5UnPadding(plainText, blockSize)
}
