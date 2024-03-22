/**
 * Author: anarckk anarckk@gamil.com
 * Date: 2023-06-26 10:53:39
 * LastEditTime: 2023-06-26 13:47:00
 * Description: 这是2022-05-17写的一个程序，想要打印文件夹的结构，已经写好一年有余了
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package beauty_print2

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitea.bee.anarckk.me/anarckk/go_util/go_unit"
)

type File struct {
	Parent   *File
	AbsPath  string
	Name     string
	IsDir    bool
	Size     int64
	Children *[]File
}

func (file *File) TotalChildSize() int64 {
	if file.Children == nil {
		return 0
	}
	var result int64
	for i := 0; i < len(*file.Children); i++ {
		result += (*file.Children)[i].Size
		result += (*file.Children)[i].TotalChildSize()
	}
	return result
}

func (file *File) BeautyString() string {
	t := file.Size + file.TotalChildSize()
	warn := ""
	if t > 500*1024 {
		warn = "大于500k"
	}
	return fmt.Sprintf("%s %s %s", file.Name, go_unit.HumanReadableByteCountBin(t), warn)
}

// GetFile 获得文件
func GetFile(path string) (*File, error) {
	var _file File
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
func ListFile(path string, parent *File) (*[]File, error) {
	fileInfoList, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var fileList []File
	for i := 0; i < len(fileInfoList); i++ {
		_fileInfo := &fileInfoList[i]
		var _file File
		if (*_fileInfo).IsDir() {
			_children, err := ListFile(path+(*_fileInfo).Name()+"/", &_file)
			if err == nil {
				_file.Children = _children
			}
			_file.AbsPath = path + (*_fileInfo).Name() + "/"
		} else {
			_file.Children = nil
			_file.AbsPath = path + (*_fileInfo).Name()
			_file.Size = (*_fileInfo).Size()
		}
		_file.Name = (*_fileInfo).Name()
		_file.IsDir = (*_fileInfo).IsDir()
		_file.Parent = parent
		fileList = append(fileList, _file)
	}
	return &fileList, nil
}

// BeautyPrint 以优美的格式输出文件夹目录结构
func BeautyPrint(_file *File, prefix string, isLast bool) {
	if _file == nil {
		return
	}
	if _file.IsDir {
		fmt.Println(prefix + "|--> " + _file.BeautyString())
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
		fmt.Println(prefix + "|--> " + _file.BeautyString())
	}
}

func PrintFileList(fileList *[]File, parent string) {
	if fileList == nil {
		return
	}
	for i := 0; i < len(*fileList); i++ {
		if (*fileList)[i].IsDir {
			PrintFileList((*fileList)[i].Children, (*fileList)[i].AbsPath)
		} else {
			fmt.Println((*fileList)[i].AbsPath)
		}
	}
}
