package main

import (
	"log"

	"gitea.bee.anarckk.me/anarckk/go_util/go_code"
)

func main() {
	md5, _ := go_code.GetFileMD5("/tmp/random-test-file-java")
	log.Println(md5)
}
