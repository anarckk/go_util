package go_aes

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"

	"gitea.bee.anarckk.me/anarckk/go_util/go_code"
	"gitea.bee.anarckk.me/anarckk/go_util/go_file"
)

func EncryptCtrOutputStream(outputStream io.Writer, key string, iv string) (io.Writer, error) {
	keyBytes, err := go_code.Base64Str2Bytes(key)
	if err != nil {
		return nil, err
	}
	ivBytes, err := go_code.Base64Str2Bytes(iv)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	outputStream.Write(ivBytes)
	stream := cipher.NewCTR(block, ivBytes)
	return &cipher.StreamWriter{S: stream, W: outputStream}, nil
}

func DecryptCtrInputStream(inputStream io.Reader, key string) (io.Reader, error) {
	ivBytes, err := go_file.ReadNBytesFromInputStream(inputStream, 16)
	if err != nil {
		return nil, err
	}
	keyBytes, err := go_code.Base64Str2Bytes(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, ivBytes)
	return &cipher.StreamReader{S: stream, R: inputStream}, nil
}

func EncryptCtrFile(src string, encryptedAbsPath string, key string, iv string) error {
	inputStream, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inputStream.Close()
	bis := bufio.NewReader(inputStream)
	outputStream, err := os.Create(encryptedAbsPath)
	if err != nil {
		return err
	}
	defer outputStream.Close()
	bos := bufio.NewWriter(outputStream)
	defer bos.Flush()

	cos, err := EncryptCtrOutputStream(bos, key, iv)
	if err != nil {
		return err
	}
	_, err = io.Copy(cos, bis)
	return err
}

func DecryptCtrFile(encryptedAbsPath string, decryptedAbsPath string, key string) error {
	inputStream, err := os.Open(encryptedAbsPath)
	if err != nil {
		return err
	}
	defer inputStream.Close()
	bis := bufio.NewReader(inputStream)
	outputStream, err := os.Create(decryptedAbsPath)
	if err != nil {
		return err
	}
	defer outputStream.Close()
	bos := bufio.NewWriter(outputStream)
	defer bos.Flush()

	cis, err := DecryptCtrInputStream(bis, key)
	if err != nil {
		return err
	}
	_, err = io.Copy(bos, cis)
	return err
}
