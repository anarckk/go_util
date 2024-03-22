/**
 * Author: anarckk anarckk@gamil.com
 * Date: 2023-06-26 10:34:40
 * LastEditTime: 2023-06-26 10:51:00
 * Description:
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package beauty_print

import (
	"fmt"
	"os"
)

type MyFile struct {
	Parent   *MyFile
	AbsPath  string
	Name     string
	IsDir    bool
	Size     int64
	Children *[]MyFile
}

func (file *MyFile) SimpleString() string {
	if file.IsDir {
		return file.Name + "/"
	} else {
		return file.Name
	}
}

// GetFile 获得文件
func GetFile(path string) (*MyFile, error) {
	var _file MyFile
	file, err := os.Open(path)
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		return nil, err
	}
	_fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	_children, err := ListFile(path, &_file)
	if err != nil {
		return nil, err
	}

	_file.AbsPath = path
	_file.Name = _fileInfo.Name()
	_file.IsDir = _fileInfo.IsDir()
	_file.Children = _children
	return &_file, nil
}

// ListFile 迭代便利整个路径下的所有文件夹
// 注意，path的最后一个字符，必须有 '/'
func ListFile(path string, parent *MyFile) (*[]MyFile, error) {
	var char string = path[len(path)-1:]
	if !(char == "/" || char == "\\") {
		return nil, fmt.Errorf("last char must \"/\" or \"\\\", but is %s", path)
	}
	fileEntryList, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var fileList []MyFile
	for i := 0; i < len(fileEntryList); i++ {
		fileEntry := &fileEntryList[i]
		var myFile MyFile
		if (*fileEntry).IsDir() {
			_children, err := ListFile(path+(*fileEntry).Name()+"/", &myFile)
			if err == nil {
				myFile.Children = _children
			}
			myFile.AbsPath = path + (*fileEntry).Name() + "/"
		} else {
			myFile.Children = nil
			myFile.AbsPath = path + (*fileEntry).Name()
			fileInfo, err := (*fileEntry).Info()
			if err != nil {
				return nil, err
			}
			myFile.Size = fileInfo.Size()
		}
		myFile.Name = (*fileEntry).Name()
		myFile.IsDir = (*fileEntry).IsDir()
		myFile.Parent = parent
		fileList = append(fileList, myFile)
	}
	return &fileList, nil
}

// BeautyPrint 以优美的格式输出文件夹目录结构
func BeautyPrint(_file *MyFile, prefix string, isLast bool) {
	if _file == nil {
		return
	}
	if _file.IsDir {
		fmt.Println(prefix + "|--> " + _file.SimpleString())
		if _file.Children != nil {
			for j := 0; j < len(*_file.Children); j++ {
				var _prefix string
				if isLast {
					_prefix = prefix + "   "
				} else {
					_prefix = prefix + "|  "
				}
				_isLast := j == len(*_file.Children)-1
				BeautyPrint(&(*_file.Children)[j], _prefix, _isLast)
			}
		}
	} else {
		fmt.Println(prefix + "|--> " + _file.SimpleString())
	}
}
