package main

import (
	"log"

	"gitea.bee.anarckk.me/anarckk/go_util/go_code"
)

func main() {
	// 5d41402abc4b2a76b9719d911017c592
	log.Println(go_code.GetMessageMD5("hello"))
	// 5dde1a40c4953cfde70e6c6926c8d068
	log.Println(go_code.GetFileMD5("/workdir/data2/pet-my-dev-runtime-data/runtime/workspace/io-test-space/滕王阁序.txt"))

	log.Println(go_code.Base64DecodeStr("aGVsbG8gd29ybGQ="))
}
