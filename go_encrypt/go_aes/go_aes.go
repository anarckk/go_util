package go_aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"

	"gitea.bee.anarckk.me/anarckk/go_util/go_code"
	aes2 "github.com/wumansgy/goEncrypt/aes"
)

func GenAesKey() (string, error) {
	key := make([]byte, 32) // 256位密钥
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return go_code.Base64Bytes2Str(key), nil
}

func EncryptEcbPkcs5Str(msg string, key string) (string, error) {
	keyBytes, err := go_code.Base64Str2Bytes(key)
	if err != nil {
		return "", err
	}
	return aes2.AesEcbEncryptBase64([]byte(msg), keyBytes)
}

func DecryptEcbPkcs5Str(encrypted string, key string) (string, error) {
	keyBytes, err := go_code.Base64Str2Bytes(key)
	if err != nil {
		return "", err
	}
	decryptedBytes, err := aes2.AesEcbDecryptByBase64(encrypted, keyBytes)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}

// EncryptCbcPkcs5Str AES加密, CBC模式，pkcs5填充
func EncryptCbcPkcs5Str(msg string, key string, iv string) (string, error) {
	keyBytes, err := go_code.Base64Str2Bytes(key)
	if err != nil {
		return "", err
	}
	ivBytes, err := go_code.Base64Str2Bytes(iv)
	if err != nil {
		return "", err
	}
	return aes2.AesCbcEncryptBase64([]byte(msg), keyBytes, ivBytes)
}

// DecryptCbcPkcs5Str AES解密, CBC模式，pkcs5填充
func DecryptCbcPkcs5Str(encrypted string, key string, iv string) (string, error) {
	keyBytes, err := go_code.Base64Str2Bytes(key)
	if err != nil {
		return "", err
	}
	ivBytes, err := go_code.Base64Str2Bytes(iv)
	if err != nil {
		return "", err
	}
	decryptedBytes, err := aes2.AesCbcDecryptByBase64(encrypted, keyBytes, ivBytes)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// PKCS5Padding1 这个更专业一点
func PKCS5Padding1(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

// PKCS5UnPadding1 这个更专业一点
func PKCS5UnPadding1(plainText []byte, blockSize int) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length || number > blockSize {
		return nil, errors.New("padding size error please check the secret key or iv")
	}
	return plainText[:length-number], nil
}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

func EncryptCbcPkcs7Str(dataStr string, keyStr string, iv string) (string, error) {
	ivBytes, err := go_code.Base64Str2Bytes(iv)
	if err != nil {
		return "", err
	}
	data := []byte(dataStr)
	key, err := go_code.Base64Str2Bytes(keyStr)
	if err != nil {
		return "", err
	}

	// NewCipher creates and returns a new cipher.Block. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, ivBytes)
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return go_code.Base64Bytes2Str(crypted), nil
}

func DecryptCbcPkcs7Str(encrypted string, keyStr string, iv string) (string, error) {
	ivBytes, err := go_code.Base64Str2Bytes(iv)
	if err != nil {
		return "", err
	}
	key, err := go_code.Base64Str2Bytes(keyStr)
	if err != nil {
		return "", err
	}
	data, err := go_code.Base64Str2Bytes(encrypted)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, ivBytes)
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return "", err
	}
	return string(crypted), nil
}

func EncryptCbcPkcs5Str2(dataStr string, keyStr string, iv string) (string, error) {
	ivBytes, err := go_code.Base64Str2Bytes(iv)
	if err != nil {
		return "", err
	}
	data := []byte(dataStr)
	key, err := go_code.Base64Str2Bytes(keyStr)
	if err != nil {
		return "", err
	}

	// NewCipher creates and returns a new cipher.Block. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//填充
	encryptBytes := PKCS5Padding1(data, block.BlockSize())
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, ivBytes)
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return go_code.Base64Bytes2Str(crypted), nil
}

func DecryptCbcPkcs5Str2(encrypted string, keyStr string, iv string) (string, error) {
	ivBytes, err := go_code.Base64Str2Bytes(iv)
	if err != nil {
		return "", err
	}
	key, err := go_code.Base64Str2Bytes(keyStr)
	if err != nil {
		return "", err
	}
	data, err := go_code.Base64Str2Bytes(encrypted)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, ivBytes)
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去填充
	crypted, err = PKCS5UnPadding1(crypted, blockMode.BlockSize())
	if err != nil {
		return "", err
	}
	return string(crypted), nil
}
