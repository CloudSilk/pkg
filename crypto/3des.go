package crypto

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

// TripleDesEncrypt 3DES加密
func TripleDesEncrypt(plainText, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
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

// TripleDesDecrypt 3DES解密
func TripleDesDecrypt(cipherText, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
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
