/*
@Time   : 2020/11/25 16:09
@Author : Haichao Wang
*/
package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"io"
)

//RSA公钥私钥产生
func GenerateRsaKey() (err error, prvkey, pubkey string) {
	var prvkeyBytes, pubkeyBytes []byte
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err, "", ""
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkeyBytes = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err, "", ""
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkeyBytes = pem.EncodeToMemory(block)
	prvkey = string(prvkeyBytes)
	pubkey = string(pubkeyBytes)
	return
}

// 生成密钥对
func CreateKeys(publicKeyWriter, privateKeyWriter io.Writer, keyLength int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return err
	}
	derStream := MarshalPKCS8PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	err = pem.Encode(privateKeyWriter, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	err = pem.Encode(publicKeyWriter, block)
	if err != nil {
		return err
	}
	return nil
}

func MarshalPKCS8PrivateKey(key *rsa.PrivateKey) []byte {
	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}
	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = x509.MarshalPKCS1PrivateKey(key)
	k, _ := asn1.Marshal(info)
	return k
}
