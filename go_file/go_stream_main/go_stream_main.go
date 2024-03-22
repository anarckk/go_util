package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"gitea.bee.anarckk.me/anarckk/go_util/go_bit"
	"gitea.bee.anarckk.me/anarckk/go_util/go_code"
	"gitea.bee.anarckk.me/anarckk/go_util/go_file"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func handleInputStream(inputStream io.ReadWriteCloser, folder string) {
	defer inputStream.Close()
	name, err := go_file.BitReadStringFromStream(inputStream)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("正在拉取文件: " + name)
	md5, err := go_file.BitReadStringFromStream(inputStream)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("    md5: " + md5)
	absPath := folder + "/" + name
	_, err = go_file.ReliableFromSocketToFile(inputStream, absPath)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("服务器正在计算md5: " + name)
	nextMD5, err := go_code.GetFileMD5(absPath)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("nextMD5: " + nextMD5)
	if md5 == nextMD5 {
		log.Println("成功拉取文件: " + name)
	} else {
		log.Println("拉取失败,md5不一致")
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func server(port int, folder string) error {
	address := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	log.Println("listening: " + address)
	log.Println("folder: " + folder)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("收到客户端连接")
		handleInputStream(conn, folder)
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func send(server string, port int, absPath string) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		return err
	}
	defer conn.Close()
	name := go_file.GetFileName(absPath)
	log.Println("正在发送文件: " + name)
	err = go_file.BitWriteStringToOutputStream(name, conn)
	if err != nil {
		return err
	}
	md5, err := go_code.GetFileMD5(absPath)
	if err != nil {
		return err
	}
	log.Println("文件md5: " + md5)
	err = go_file.BitWriteStringToOutputStream(md5, conn)
	if err != nil {
		return err
	}
	err = go_file.ReliableFromFileToSocket(absPath, conn)
	if err != nil {
		return err
	}
	log.Println("成功发送文件: " + absPath)
	time.Sleep(100 * time.Hour)
	return nil
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testServer() {
	go func() {
		err := server(7782, "/pet-workdir/data2/pet-dev-var/var/workspace/io-test-space/go-transfer")
		if err != nil {
			log.Println(err)
		}
	}()

	err := send("localhost", 7782, "/pet-workdir/data2/pet-dev-var/var/workspace/io-test-space/vscode-server.tar")
	if err != nil {
		log.Println(err)
	}

	time.Sleep(100 * time.Hour)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func test1() {
	is := go_file.NewRandomReader(100)
	byteArray, _ := go_file.FromInputStreamToByteArray(is)
	log.Println(go_bit.BytesToBytesStr(byteArray))
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testCopyPeek() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	randomReader := go_file.NewRandomReader(5 * 1024 * 1024 * 1024)
	nullWriter := go_file.NewNullWriter()
	go_file.NewStreamCopy().Source(randomReader).Peek(go_file.AutoPrintSpeed(ctx)).Target(nullWriter).RunCopy()
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testSocketFileTransfer() {
	fileAbsPath := "/tmp/random-test-file"
	// fileMd5 := "53d8527e7912c0a176433be82de49822"
	fileAbsPath2 := "/tmp/random-test-file2"

	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		defer wait.Done()
		listener, err := net.Listen("tcp", ":18888")
		if err != nil {
			panic(err)
		}
		defer listener.Close()
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		log.Println("server: 收到客户端连接")
		go_file.NewBitFromFileToOutputStream().AbsPath(fileAbsPath).Writer(conn).ShowProgress(true).Run()
		if err != nil {
			panic(err)
		}
		log.Println("server: 发送文件完成")
	}()
	time.Sleep(time.Second)
	go func() {
		defer wait.Done()
		conn, err := net.Dial("tcp", "localhost:18888")
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		log.Println("client: 已连接服务端")
		err = go_file.NewBitFromInputStreamToFile().Reader(conn).AbsPath(fileAbsPath2).Run()
		if err != nil {
			panic(err)
		}
	}()
	wait.Wait()
	log.Println("测试结束")
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testAutoSpeed() {
	var len uint64 = 5 * 1024 * 1024 * 1024
	randomReader := go_file.NewRandomReader(len)
	nullWriter := go_file.NewNullWriter()
	ais := go_file.NewAutoSpeedInputStream(randomReader, len)
	defer ais.Close()
	go_file.Copy(nullWriter, ais)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	testCopyPeek()
}
