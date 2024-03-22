package main

import (
	"log"
	"net"

	"gitea.bee.anarckk.me/anarckk/go_util/go_file"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func getStupid() string {
	str, _ := go_file.ReadFileToString("/pet-workdir/data2/pet/pet-dev/pet-workspace/io-test-space/stupid/stupid.txt")
	return str
}

// 启一个服务器来测试java
func TestJava() {
	fileAbsPath := "/tmp/random-test-file"
	// fileMd5 := "53d8527e7912c0a176433be82de49822"
	listener, err := net.Listen("tcp", ":18888")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Println("等待连接...")
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("server: 收到客户端连接")
	go_file.NewBitFromFileToOutputStream().AbsPath(fileAbsPath).Writer(conn).Run()
	log.Println("server: 发送文件完成")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
