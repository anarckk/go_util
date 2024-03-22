package go_file

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Deprecated: 废弃
// BitFromFileToOutputStream 按照自己规则(传输指定的字节量) - 把文件写入到流中
// 仅支持同一时间单向传播
// 传大文件，可靠
func ReliableFromFileToSocket(absPath string, socket io.ReadWriter) error {
	reliableNum := 50 * 1024 * 1024

	socketBos := bufio.NewWriter(socket)
	defer socketBos.Flush()

	err := BitWriteInt64ToOutputStream(GetFileSize(absPath), socketBos)
	if err != nil {
		return err
	}

	fis, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer fis.Close()
	bis := bufio.NewReader(fis)
	var tmp = make([]byte, 1024)
	var tmpTotal = 0 // 本次的 reliableNum 范围的数量
	var tmpReadLen = 0
	for {
		tmpReadLen, err = bis.Read(tmp[:])
		if err == io.EOF {
			socketBos.Flush()
			_, err := ReadNBytesFromInputStream(socket, 1)
			if err != nil {
				return err
			}
			break
		}
		if err != nil {
			return err
		}
		_, err = socketBos.Write(tmp[0:tmpReadLen])
		if err != nil {
			return err
		}
		tmpTotal += tmpReadLen
		if tmpTotal >= reliableNum {
			socketBos.Flush()
			_, err := ReadNBytesFromInputStream(socket, 1)
			if err != nil {
				return err
			}
			tmpTotal = 0
			tmp = make([]byte, 1024)
		} else if reliableNum-tmpTotal < 1024 {
			tmp = make([]byte, reliableNum-tmpTotal)
		}
	}
	return nil
}

// Deprecated: 废弃
// ReadIntFromInputStream 按照自己规则(传输指定的字节量) - 从流中读取指定字节，写入到文件中
// 仅支持同一时间单向传播
// 传大文件，可靠
// 传10G文件，可靠，从理论上将，应该多大都可靠
// 最后返回文件的md5
func ReliableFromSocketToFile(socket io.ReadWriter, absPath string) (string, error) {
	reliableNum := 50 * 1024 * 1024

	socketBis := bufio.NewReader(socket)

	fos, err := os.Create(absPath)
	if err != nil {
		return "", err
	}
	defer fos.Close()
	bos := bufio.NewWriter(fos)
	defer bos.Flush()

	fileSize, err := BitReadInt64FromInputStream(socketBis)
	if err != nil {
		return "", err
	}

	var total int64 = 0 // 已经读取的总数量
	var tmpTotal = 0    // 本次的 reliableNum 范围的数量
	var tmpReadLen = 0
	var tmp = make([]byte, 1024)

	md5h := md5.New()
	for {
		tmpReadLen, err = socketBis.Read(tmp)
		if err != nil {
			return "", err
		}
		_, err = bos.Write(tmp[0:tmpReadLen])
		if err != nil {
			return "", err
		}
		_, err = md5h.Write(tmp[0:tmpReadLen])
		if err != nil {
			return "", err
		}
		total += int64(tmpReadLen)
		if total >= fileSize {
			bos.Flush()
			_, err := socket.Write([]byte{0x00})
			if err != nil {
				return "", err
			}
			break
		}
		tmpTotal += tmpReadLen
		if tmpTotal >= reliableNum {
			bos.Flush()
			_, err := socket.Write([]byte{0x00})
			if err != nil {
				return "", err
			}
			tmpTotal = 0
		}
	}
	return hex.EncodeToString(md5h.Sum(nil)), nil
}
