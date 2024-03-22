/**
 * Author: anarckk anarckk@gmail.com
 * Date: 2023-04-29 09:33:46
 * LastEditTime: 2023-04-29 22:33:36
 * Description:
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package go_file

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// GetFileName 返回文件名称，包括文件后缀名
//
//	@param path "/root/速度计算ai提示词.txt"
//	@return string 速度计算ai提示词.txt
func GetFileName(path string) string {
	return GetFileName1(path)
}

func GetFileName1(path string) string {
	// 使用strings包中的Split函数将路径按照"/"分割成一个字符串切片
	segments := strings.Split(path, "/")
	// 获取切片中最后一个元素，即文件名称
	fileName := segments[len(segments)-1]
	// 返回文件名称
	return fileName
}

func GetFileName2(path string) string {
	stat, err := os.Stat(path)
	if err != nil {
		return GetFileName1(path)
	}
	return stat.Name()
}

func FileExists(absPath string) bool {
	// 使用os.Stat函数获取文件信息
	_, err := os.Stat(absPath)
	return !os.IsNotExist(err)
}

func Mkdir(absPath string) error {
	return os.Mkdir(absPath, os.ModePerm)
}

func MkdirAll(absPath string) error {
	return os.MkdirAll(absPath, os.ModePerm)
}

// 这个api遵循这样一种原则
// 如果你要用我，请上层自己确保输入的路径的确是一个路径，否则，我的行为将是无法预料的
func GetFileSize(absPath string) int64 {
	stat, _ := os.Stat(absPath)
	return stat.Size()
}

// GetParentAbsPath 获取文件父目录的绝对路径
//
//	@param absPath
//	@return string
func GetParentAbsPath(absPath string) (string, error) {
	parentDir := filepath.Dir(absPath)
	absParentDir, err := filepath.Abs(parentDir)
	if err != nil {
		return "", err
	}
	return absParentDir, nil
}

func IsDir(absPath string) (bool, error) {
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

func IsFile(absPath string) (bool, error) {
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return false, err
	}
	return !fileInfo.IsDir(), nil
}

// Deprecated: 废弃
func ListDir(absPath string) ([]string, error) {
	return WalkDir(absPath)
}

func WalkDir(absPath string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(absPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ReadFileToString(fileAbsPath string) (string, error) {
	_bytes, err := ReadFileToByteArray(fileAbsPath)
	if err != nil {
		return "", err
	}
	return string(_bytes), nil
}

func WriteStrToFile(str string, fileAbsPath string) error {
	return WriteByteArrayToFile([]byte(str), fileAbsPath)
}

// ReadFileToByteArray 从文件中读取文件内容，如果文件的父文件夹不存在，则将报错
//
//	@param fileAbsPath 待读取的文件的绝对路径
//	@return []byte 文件的内容
//	@return error 是否出错
func ReadFileToByteArray(fileAbsPath string) ([]byte, error) {
	return ReadFileToByteArray2(fileAbsPath)
}

func ReadFileToByteArray1(fileAbsPath string) ([]byte, error) {
	open, err := os.Open(fileAbsPath)
	if err != nil {
		return nil, err
	}
	defer open.Close()
	var bytes [1024]byte
	var result []byte
	for {
		n, err := open.Read(bytes[:])
		if err == io.EOF {
			return result, nil
		}
		if err != nil {
			return nil, err
		}
		result = append(result, bytes[:n]...)
	}
}

func ReadFileToByteArray2(fileAbsPath string) ([]byte, error) {
	open, err := os.Open(fileAbsPath)
	if err != nil {
		return nil, err
	}
	defer open.Close()
	bytes, err := io.ReadAll(open)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func WriteByteArrayToFile(bytes []byte, absPath string) error {
	return WriteByteArrayToFile2(bytes, absPath)
}

func WriteByteArrayToFile1(bytes []byte, absPath string) error {
	return os.WriteFile(absPath, bytes, 0666)
}

func WriteByteArrayToFile2(bytes []byte, absPath string) error {
	f, err := os.Create(absPath)
	if err != nil {
		return err
	}
	_, err = f.Write(bytes)
	return err
}

func CopyFile(src string, dst string) error {
	return CopyFile2(src, dst)
}

func CopyFile1(src string, dst string) error {
	inputStream, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inputStream.Close()
	outputStream, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outputStream.Close()
	var bytes [1024]byte
	for {
		n, err := inputStream.Read(bytes[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		outputStream.Write(bytes[:n])
	}
	return nil
}

func CopyFile2(src string, dst string) error {
	inputStream, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inputStream.Close()
	outputStream, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outputStream.Close()
	_, err = io.Copy(outputStream, inputStream)
	return err
}

func MoveFile(src string, dst string) {
	os.Rename(src, dst)
}

func DelFile(absPath string) error {
	return os.Remove(absPath)
}

// GenerateTestFile 在指定位置创建一个指定大小的测试文件，并返回该文件的md5
func GenerateTestFile(absPath string, size int64) (string, error) {
	outputStream, err := os.Create(absPath)
	if err != nil {
		return "", err
	}
	bos := bufio.NewWriter(outputStream)
	md5h := md5.New()
	writer := io.MultiWriter(bos, md5h)
	randReader := NewRandomReader(uint64(size))
	io.TeeReader(randReader, bos)
	io.Copy(writer, randReader)
	return hex.EncodeToString(md5h.Sum(nil)), nil
}
