package crypto

import (
	"testing"
)

func TestDes(t *testing.T) {
	// 8字节的密钥
	key := []byte("12345678")
	// 明文
	plainText := []byte("Hello, DES!")

	t.Logf("原文: %s\n", plainText)

	// 加密
	ciphertext, err := DesEncrypt(plainText, key)
	if err != nil {
		t.Fatalf("加密错误：%v\n", err)
	}
	t.Logf("密文: %x\n", ciphertext)

	// 解密
	decryptedText, err := DesDecrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("解密错误：%v\n", err)
	}
	t.Logf("解密后的原文: %s\n", decryptedText)
}

// TestDesEncryptDecrypt 测试DES加解密
func TestDesEncryptDecrypt(t *testing.T) {
	key := []byte("12345678")
	plainText := []byte("Hello, DES!")

	// 正常情况
	ciphertext, err := DesEncrypt(plainText, key)
	if err != nil {
		t.Fatalf("DesEncrypt error: %v", err)
	}

	decryptedText, err := DesDecrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("DesDecrypt error: %v", err)
	}

	if string(decryptedText) != string(plainText) {
		t.Fatalf("Decrypted text does not match original. Got %s, wanted %s", decryptedText, plainText)
	}

	// 错误的密钥长度
	_, err = DesEncrypt(plainText, []byte("short"))
	if err == nil {
		t.Fatalf("Expected an error for short key")
	}
}

func TestPKCS5UnPadding(t *testing.T) {
	// 正常情况
	validPaddedData := []byte("Hello world\x03\x03\x03")
	_, err := PKCS5UnPadding(validPaddedData, 8)
	if err != nil {
		t.Fatalf("PKCS5UnPadding error: %v", err)
	}

	// 异常情况：错误的填充大小
	invalidPaddedData1 := []byte("Hello world\x05\x05\x05")
	_, err = PKCS5UnPadding(invalidPaddedData1, 8)
	if err == nil {
		t.Fatalf("Expected an error for invalid padding size")
	}

	// 异常情况：填充字节不一致
	invalidPaddedData2 := []byte("Hello world\x03\x03\x02")
	_, err = PKCS5UnPadding(invalidPaddedData2, 8)
	if err == nil {
		t.Fatalf("Expected an error for inconsistent padding bytes")
	}

	// 异常情况：空的密文
	_, err = PKCS5UnPadding([]byte(""), 8)
	if err == nil {
		t.Fatalf("Expected an error for empty cipherText")
	}
}
