package go_rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"

	"gitea.bee.anarckk.me/anarckk/go_util/go_code"
)

// GenerateKeyPair generates a new key pair
func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privkey, &privkey.PublicKey, nil
}

// PrivateKeyToBytes private key to bytes
func PrivateKeyToBytes(priv *rsa.PrivateKey) ([]byte, error) {
	bytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, err
	}
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: bytes,
		},
	)
	return privBytes, nil
}

// PublicKeyToBytes public key to bytes
func PublicKeyToBytes(pub *rsa.PublicKey) ([]byte, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return pubASN1, err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes, nil
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	return PKCS8BytesToPrivateKey(b)
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not ok")
	}
	return key, nil
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	hash := sha512.New()
	return rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	hash := sha512.New()
	return rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
}

func PKCS8PrivateKeyToBytes(priv *rsa.PrivateKey) ([]byte, error) {
	return x509.MarshalPKCS8PrivateKey(priv)
}
func PKCS8PublicKeyToBytes(pub *rsa.PublicKey) ([]byte, error) {
	return x509.MarshalPKIXPublicKey(pub)
}
func PKCS8BytesToPrivateKey(raw []byte) (*rsa.PrivateKey, error) {
	k, err := x509.ParsePKCS8PrivateKey(raw)
	if priK, ok := k.(*rsa.PrivateKey); ok {
		return priK, err
	}
	return nil, errors.New("类型转换失败")
}
func PKCS8BytesToPublicKey(raw []byte) (*rsa.PublicKey, error) {
	ifc, err := x509.ParsePKIXPublicKey(raw)
	if err != nil {
		return nil, err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not ok")
	}
	return key, nil
}

// GenerateKeyPairStr 生成公私钥
//
//	@return string 私钥字符串
//	@return string 公钥字符串
func GenerateKeyPairStr() (string, string, error) {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}
	priBytes, err := PKCS8PrivateKeyToBytes(privkey)
	if err != nil {
		return "", "", err
	}
	pubBytes, err := PKCS8PublicKeyToBytes(&privkey.PublicKey)
	if err != nil {
		return "", "", err
	}
	return go_code.Base64Encode(priBytes), go_code.Base64Encode(pubBytes), nil
}

// EncryptByPubKeyStrOAEP 使用公钥加密
//
//	@param msg 明文消息
//	@param pub 公钥字符串
//	@return string 密文消息
func EncryptByPubKeyStrOAEP(msg string, pub string) (string, error) {
	pubBytes, err := go_code.Base64Decode(pub)
	if err != nil {
		return "", err
	}
	pubK, err := PKCS8BytesToPublicKey(pubBytes)
	if err != nil {
		return "", err
	}

	hash := sha512.New()
	encryptedBytes, err := rsa.EncryptOAEP(hash, rand.Reader, pubK, []byte(msg), nil)
	if err != nil {
		return "", err
	}
	return go_code.Base64Encode(encryptedBytes), err
}

// DecryptByPriKeyStrOAEP 用私钥解密
//
//	@param encryptedMsg 密文
//	@param pri 私钥字符串
//	@return string 明文消息
func DecryptByPriKeyStrOAEP(encryptedMsg string, pri string) (string, error) {
	priBytes, err := go_code.Base64Decode(pri)
	if err != nil {
		return "", err
	}
	priK, err := PKCS8BytesToPrivateKey(priBytes)
	if err != nil {
		return "", err
	}
	_encryptedMsgBytes, err := go_code.Base64Decode(encryptedMsg)
	if err != nil {
		return "", err
	}

	hash := sha512.New()
	decryptedBytes, err := rsa.DecryptOAEP(hash, rand.Reader, priK, _encryptedMsgBytes, nil)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}

// EncryptByPubKeyStr 使用公钥加密
//
//	@param msg 明文消息
//	@param pub 公钥字符串
//	@return string 密文消息
func EncryptByPubKeyStr(msg string, pub string) (string, error) {
	pubBytes, err := go_code.Base64Decode(pub)
	if err != nil {
		return "", err
	}
	pubK, err := PKCS8BytesToPublicKey(pubBytes)
	if err != nil {
		return "", err
	}

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubK, []byte(msg))
	if err != nil {
		return "", err
	}
	return go_code.Base64Encode(encryptedBytes), err
}

// DecryptByPriKeyStr 用私钥解密
//
//	@param encryptedMsg 密文
//	@param pri 私钥字符串
//	@return string 明文消息
func DecryptByPriKeyStr(encryptedMsg string, pri string) (string, error) {
	priBytes, err := go_code.Base64Decode(pri)
	if err != nil {
		return "", err
	}
	priK, err := PKCS8BytesToPrivateKey(priBytes)
	if err != nil {
		return "", err
	}
	encryptedMsgBytes, err := go_code.Base64Decode(encryptedMsg)
	if err != nil {
		return "", err
	}
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priK, encryptedMsgBytes)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}
