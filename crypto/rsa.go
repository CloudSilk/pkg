package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("生成密钥对失败", err)
		return nil, nil, err
	}
	return privkey, &privkey.PublicKey, nil
}

func RSAEncrypt(plaintext []byte, pubkey *rsa.PublicKey) ([]byte, error) {
	// 使用PKCS1v15进行加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, plaintext)
	if err != nil {
		fmt.Println("加密失败", err)
		return nil, err
	}
	return ciphertext, nil
}

func RSADecrypt(ciphertext []byte, privkey *rsa.PrivateKey) ([]byte, error) {
	// 使用PKCS1v15进行解密
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privkey, ciphertext)
	if err != nil {
		fmt.Println("解密失败", err)
		return nil, err
	}
	return plaintext, nil
}

func PrivateKeyToString(priv *rsa.PrivateKey) string {
	privBytes := x509.MarshalPKCS1PrivateKey(priv)
	privPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privBytes,
		},
	)
	return string(privPem)
}

func PublicKeyToString(pub *rsa.PublicKey) (string, error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		fmt.Println("公钥转换失败:", err)
		return "", err
	}
	pubPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubBytes,
		},
	)
	return string(pubPem), nil
}

func StringToPrivateKey(privStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privStr))
	if block == nil {
		return nil, fmt.Errorf("无法解码私钥字符串")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("无法解析私钥: %v", err)
	}
	return priv, nil
}

func StringToPublicKey(pubStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubStr))
	if block == nil {
		return nil, fmt.Errorf("无法解码公钥字符串")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("无法解析公钥: %v", err)
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("不是有效的RSA公钥")
	}
	return rsaPub, nil
}
