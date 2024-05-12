package crypto

import (
	"bytes"
	"testing"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	_, _, err := GenerateRSAKeyPair()
	if err != nil {
		t.Errorf("生成密钥对失败: %v", err)
	}
}

func TestRSAEncryptDecrypt(t *testing.T) {
	priv, pub, err := GenerateRSAKeyPair()
	if err != nil {
		t.Errorf("生成密钥对失败: %v", err)
		return
	}

	plaintext := []byte("测试明文")
	ciphertext, err := RSAEncrypt(plaintext, pub)
	if err != nil {
		t.Errorf("加密失败: %v", err)
		return
	}

	decryptedText, err := RSADecrypt(ciphertext, priv)
	if err != nil {
		t.Errorf("解密失败: %v", err)
		return
	}

	if !bytes.Equal(plaintext, decryptedText) {
		t.Errorf("解密文本与原文不匹配")
	}
}

func TestPrivateKeyToString(t *testing.T) {
	priv, _, err := GenerateRSAKeyPair()
	if err != nil {
		t.Errorf("生成密钥对失败: %v", err)
		return
	}

	privStr := PrivateKeyToString(priv)
	if privStr == "" {
		t.Errorf("私钥转字符串失败")
	}
}

func TestPublicKeyToString(t *testing.T) {
	_, pub, err := GenerateRSAKeyPair()
	if err != nil {
		t.Errorf("生成密钥对失败: %v", err)
		return
	}

	pubStr, err := PublicKeyToString(pub)
	if err != nil {
		t.Errorf("公钥转字符串失败: %v", err)
	}
	if pubStr == "" {
		t.Errorf("公钥转字符串失败")
	}
}

func TestStringToPrivateKey(t *testing.T) {
	priv, _, err := GenerateRSAKeyPair()
	if err != nil {
		t.Errorf("生成密钥对失败: %v", err)
		return
	}

	privStr := PrivateKeyToString(priv)
	priv2, err := StringToPrivateKey(privStr)
	if err != nil {
		t.Errorf("字符串转私钥失败: %v", err)
	}
	if priv2 == nil {
		t.Errorf("字符串转私钥失败")
	}
}

func TestStringToPublicKey(t *testing.T) {
	_, pub, err := GenerateRSAKeyPair()
	if err != nil {
		t.Errorf("生成密钥对失败: %v", err)
		return
	}

	pubStr, err := PublicKeyToString(pub)
	if err != nil {
		t.Errorf("公钥转字符串失败: %v", err)
	}
	pub2, err := StringToPublicKey(pubStr)
	if err != nil {
		t.Errorf("字符串转公钥失败: %v", err)
	}
	if pub2 == nil {
		t.Errorf("字符串转公钥失败")
	}
}
