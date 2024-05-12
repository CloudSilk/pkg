package crypto

import (
	"testing"
)

func TestAES(t *testing.T) {
	// 16, 24, 32字节的密钥对应AES-128, AES-192, AES-256
	key := []byte("1234567890123456")
	// 明文
	plainText := []byte("Hello, AES!")

	t.Logf("原文: %s\n", plainText)

	// 加密
	ciphertext, err := AesEncrypt(plainText, key)
	if err != nil {
		t.Fatalf("加密错误：%v\n", err)
		return
	}
	t.Logf("密文: %x\n", ciphertext)

	// 解密
	decryptedText, err := AesDecrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("解密错误：%v\n", err)
		return
	}
	t.Logf("解密后的原文: %s\n", decryptedText)
}

// TestAesEncryptDecrypt 测试AES加解密
func TestAesEncryptDecrypt(t *testing.T) {
	key := []byte("1234567890123456")
	plainText := []byte("Hello, AES!")

	// 正常情况
	ciphertext, err := AesEncrypt(plainText, key)
	if err != nil {
		t.Fatalf("AesEncrypt error: %v", err)
	}

	decryptedText, err := AesDecrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("AesDecrypt error: %v", err)
	}

	if string(decryptedText) != string(plainText) {
		t.Fatalf("Decrypted text does not match original. Got %s, wanted %s", decryptedText, plainText)
	}

	// 错误的密钥长度
	_, err = AesEncrypt(plainText, []byte("shortkey"))
	if err == nil {
		t.Fatalf("Expected an error for short key")
	}

	// 空的明文
	_, err = AesEncrypt([]byte(""), key)
	if err != nil {
		t.Fatalf("AesEncrypt error with empty plaintext: %v", err)
	}

	// 空的密文
	_, err = AesDecrypt([]byte(""), key)
	if err == nil {
		t.Fatalf("Expected an error for empty ciphertext")
	}
}
