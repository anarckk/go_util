package go_code

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

func GetMessageMD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) // 将[]byte转成16进制
	return md5str
}

// GetFileMD5 读取文件的md5
//
//	@param absPath 绝对路径
func GetFileMD5(absPath string) (string, error) {
	// 文件全路径名
	pFile, err := os.Open(absPath)
	if err != nil {
		return "", err
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)
	return hex.EncodeToString(md5h.Sum(nil)), nil
}

var _ IGetBinMd5 = &getBinMd5{}

type IGetBinMd5 interface {
	Write(_bytes []byte) (int, error)

	GetMd5() string
}

type getBinMd5 struct {
	h hash.Hash
}

func (g *getBinMd5) Write(_bytes []byte) (int, error) {
	return g.h.Write(_bytes)
}

func (g *getBinMd5) GetMd5() string {
	return hex.EncodeToString(g.h.Sum(nil))
}

func GetBinMd5() IGetBinMd5 {
	return &getBinMd5{h: md5.New()}
}
