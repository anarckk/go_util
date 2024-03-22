package go_code

import "encoding/base64"

// Deprecated: 废弃, 改用 Base64Bytes2Str
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Deprecated: 废弃, 改用 Base64Str2Bytes
func Base64Decode(base64Str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Str)
}

func Base64Bytes2Str(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Str2Bytes(base64Str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Str)
}

func Base64EncodeStr(msg string) string {
	return Base64Bytes2Str([]byte(msg))
}

func Base64DecodeStr(encoded string) (string, error) {
	msgBytes, err := Base64Str2Bytes(encoded)
	if err != nil {
		return "", err
	}
	return string(msgBytes), nil
}
