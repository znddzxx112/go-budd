package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// rsa公钥加密
func RsaEncrypt(data, publicKey string) (string, error) {

	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", fmt.Errorf("pem.Decode(): %s", "block is empty")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509.ParsePKIXPublicKey(): %s", err.Error())
	}

	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(data))
	if err != nil {
		return "", fmt.Errorf("rsa.EncryptPKCS1v15(): %s", err.Error())
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// rsa私钥解密
func RsaDecrypt(ciphertext, privateKey string) (string, error) {

	// "hDmXeEEblgfeb+fFlzl9VMUkfFfvNB0fR7JOJBCVlOuOKMVPDStv2I0SyTQJbCkYUb5wvATjKKrewuUBLJFOjD+i3B+YHNQ2GGsVipDjf9LyjHP6SEf1eba6+6zQDxkyVGdYBhsYaywEkwloRfi1gQ/8iT4479MToBbuAQCxe08="
	keyBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64.StdEncoding.DecodeString(): %s", err.Error())
	}

	//获取私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", fmt.Errorf("pem.Decode(): %s", "block is empty")
	}

	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509.ParsePKCS1PrivateKey(): %s", err.Error())
	}

	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, keyBytes)
	if err != nil {
		return "", fmt.Errorf("rsa.DecryptPKCS1v15(): %s", err.Error())
	}
	return string(data), nil
}
