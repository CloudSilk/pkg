package crypto

import (
	"testing"
)

func Test3DES(t *testing.T) {
	// 24字节的密钥
	key := []byte("123456789012345678901234")
	// 明文
	plainText := []byte("Hello, 3DES!")

	t.Logf("原文: %s\n", plainText)

	// 加密
	ciphertext, err := TripleDesEncrypt(plainText, key)
	if err != nil {
		t.Fatalf("加密错误：%v\n", err)
		return
	}
	t.Logf("密文: %x\n", ciphertext)

	// 解密
	decryptedText, err := TripleDesDecrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("解密错误：%v\n", err)
		return
	}
	t.Logf("解密后的原文: %s\n", decryptedText)
}

// TestTripleDesEncryptDecrypt 测试3DES加解密
func TestTripleDesEncryptDecrypt(t *testing.T) {
	key := []byte("123456789012345678901234")
	plainText := []byte("Hello, 3DES!")

	// 正常情况
	ciphertext, err := TripleDesEncrypt(plainText, key)
	if err != nil {
		t.Fatalf("TripleDesEncrypt error: %v", err)
	}

	decryptedText, err := TripleDesDecrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("TripleDesDecrypt error: %v", err)
	}

	if string(decryptedText) != string(plainText) {
		t.Fatalf("Decrypted text does not match original. Got %s, wanted %s", decryptedText, plainText)
	}

	// 错误的密钥长度
	_, err = TripleDesEncrypt(plainText, []byte("shortkey"))
	if err == nil {
		t.Fatalf("Expected an error for short key")
	}
}
