package main

import (
	"fmt"
	"log"
	"time"

	"gitea.bee.anarckk.me/anarckk/go_util/go_file"
)

func Test1() {
	absPath := "/workdir/data2/pet-my-dev-runtime-data/runtime/workspace/io-test-space/滕王阁序.txt"
	log.Println(go_file.GetFileName(absPath))

	s := go_file.GetFileSize(absPath)
	log.Println(s)

	parentPath, _ := go_file.GetParentAbsPath(absPath)
	log.Println(parentPath)

	bytes, _ := go_file.ReadFileToByteArray(absPath)
	log.Println(go_file.GetFileSize(absPath))
	log.Println(len(bytes))

	nextPath, _ := go_file.GetParentAbsPath(absPath)
	nextPath += "/" + "滕王阁序copy.txt"
	// go_file.WriteContentToFile(nextPath, bytes)

	go_file.CopyFile(absPath, nextPath)
}

func Test2() {
	lists, err := go_file.WalkDir("/pet-workdir/data2/pet-dev-var/var/pet-workspace/io-test-space")
	if err != nil {
		panic(err)
	}
	for _, li := range lists {
		log.Println(li)
	}
}

func TestCreateRandomFile() {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Printf("TestCreateRandomFile use time: %s\n", end.Sub(start))
	}()
	filePath := "/tmp/random-test-file"
	md5, _ := go_file.GenerateTestFile(filePath, 1024*1024*1024*10)
	log.Println(md5) // 53d8527e7912c0a176433be82de49822
}

func main() {
	TestCreateRandomFile()
}
