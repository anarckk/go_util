package go_code

import (
	"os"
)

func Base64EncodeFile(src string, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	encoded := Base64Encode(data)
	err = os.WriteFile(dst, []byte(encoded), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Base64DecodeFile(base64File string, dst string) error {
	bytes, err := os.ReadFile(base64File)
	if err != nil {
		return err
	}
	decoded, err := Base64Decode(string(bytes))
	if err != nil {
		return err
	}
	err = os.WriteFile(dst, decoded, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
