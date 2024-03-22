package main

import (
	"log"

	"gitea.bee.anarckk.me/anarckk/go_util/go_code"
	"gitea.bee.anarckk.me/anarckk/go_util/go_encrypt/go_aes"
	"gitea.bee.anarckk.me/anarckk/go_util/go_random"
)

// test1
/* ECB
2023/11/21 23:02:35 main.go:20: key: oYfAzosupbU0hVh/mBygExUmayiHVpBpAyDk69MkExw=
2023/11/21 23:02:35 main.go:26: encrypted: UeNXs0cC0Equ8SuSSV280g==
2023/11/21 23:02:35 main.go:31: decrypted: hello world
*/
/* CBC
root@ae6e92b25de2:/pet-workdir/data2/pet-dev-var/var/pet-workspace/02.go/go_util/go_encrypt/go_aes/go_aes_main# go run .
2023/11/22 10:04:00 main.go:27: key: aYJbH+41ESY0vXusmE0hekBLg82svzEnPfYbDOCt4iQ=
2023/11/22 10:04:00 main.go:33: encrypted: CDzERx1dev9Jv3FdrKTBoA==
2023/11/22 10:04:00 main.go:38: decrypted: hello world
*/
//lint:ignore U1000 Ignore unused function temporarily for debugging
func test1() {
	iv := "zJIgKHg/wdasr6Pp5iCiLA=="
	key, err := go_aes.GenAesKey()
	if err != nil {
		panic(err)
	}
	log.Println("key: " + key)
	msg := "hello world"
	encrypted, err := go_aes.EncryptCbcPkcs5Str(msg, key, iv)
	if err != nil {
		panic(err)
	}
	log.Println("encrypted: " + encrypted)
	decrypted, err := go_aes.DecryptCbcPkcs5Str(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	log.Println("decrypted: " + decrypted)
	if decrypted != msg {
		panic(decrypted)
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func test2() {
	iv := "zJIgKHg/wdasr6Pp5iCiLA=="
	key := "cOsxbbkDQ6y8jNQVXOs70CCIi6LyPpOHXGtRNS+HdYE="
	encrypted := "jWV0rsuI7duTBVy4LCS3uA=="
	decrypted, err := go_aes.DecryptCbcPkcs5Str(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	log.Println("decrypted: " + decrypted)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func test3() {
	absPath := "/pet-workdir/data2/pet-dev-var/var/pet-workspace/io-test-space/滕王阁序.txt"
	encryptedPath := "/pet-workdir/data2/pet-dev-var/var/pet-workspace/io-test-space/go-encrypted/滕王阁序.txt.enc"
	decryptedPath := "/pet-workdir/data2/pet-dev-var/var/pet-workspace/io-test-space/go-encrypted/滕王阁序.dec.txt"
	key, err := go_aes.GenAesKey()
	if err != nil {
		panic(err)
	}
	iv := go_random.RandomBytes(16)
	err = go_aes.EncryptCtrFile(absPath, encryptedPath, key, go_code.Base64Bytes2Str(iv))
	if err != nil {
		panic(err)
	}
	err = go_aes.DecryptCtrFile(encryptedPath, decryptedPath, key)
	if err != nil {
		panic(err)
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func test4() {
	iv := go_random.RandomBytes(16)
	key, err := go_aes.GenAesKey()
	if err != nil {
		panic(err)
	}
	log.Println("key: " + key)
	msg := "hello world"
	encrypted, err := go_aes.EncryptCbcPkcs5Str2(msg, key, go_code.Base64Bytes2Str(iv))
	if err != nil {
		panic(err)
	}
	log.Println("encrypted: " + encrypted)
	decrypted, err := go_aes.DecryptCbcPkcs5Str2(encrypted, key, go_code.Base64Bytes2Str(iv))
	if err != nil {
		panic(err)
	}
	log.Println("decrypted: " + decrypted)
	if decrypted != msg {
		panic(decrypted)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	test4()
}
