package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"os"
)

const (
	privateFileName  = "private.pem"
	publicFileName   = "public.pem"
	privateKeyPrefix = "RSA PRIVATE KEY "
	publicKeyPrefix  = "RSA PUBLIC KEY "
)

var RsaEncrypt *rsaEncrypt

type rsaEncrypt struct {
	publicKey  []byte
	privateKey []byte
}

// 读取public.pem, private.pem密匙文件
func NewRsaEncrypt(publicKey, privateKey string) {
	RsaEncrypt = &rsaEncrypt{
		publicKey:  []byte(publicKey),
		privateKey: []byte(privateKey),
	}
}

func GetRsaKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateFile, err := os.Create(privateFileName)
	if err != nil {
		return err
	}
	defer privateFile.Close()
	privateBlock := pem.Block{
		Type:  privateKeyPrefix,
		Bytes: x509PrivateKey,
	}

	if err = pem.Encode(privateFile, &privateBlock); err != nil {
		return err
	}
	publicKey := privateKey.PublicKey
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	publicFile, _ := os.Create(publicFileName)
	defer publicFile.Close()
	publicBlock := pem.Block{
		Type:  publicKeyPrefix,
		Bytes: x509PublicKey,
	}
	if err = pem.Encode(publicFile, &publicBlock); err != nil {
		return err
	}
	return nil
}

func (r *rsaEncrypt) encrypt(content string) (string, error) {
	plainText := []byte(content)
	block, _ := pem.Decode(r.publicKey)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func (r *rsaEncrypt) decrypt(cryptContent string) (string, error) {
	// 私匙解密
	block, _ := pem.Decode(r.privateKey)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	cryptText, err := hex.DecodeString(cryptContent)
	if err != nil {
		return "", err
	}
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cryptText)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

// 获取可加密明文的最大长度
func (r *rsaEncrypt)GetLimitMsgSize() (int, error) {
	block, _ := pem.Decode(r.publicKey)
	if block == nil {
		return 0, errors.New("public key decode fail")
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return 0, err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	size := publicKey.Size()
	return size, nil
}

