/*
 * @Author: anarckk anarckk@gmail.com
 * @Date: 2023-12-12 09:01:31
 * @LastEditTime: 2023-12-12 10:08:29
 * @Description: https://www.cnblogs.com/xiaofengshuyu/p/5646494.html
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package go_gzip

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path"
	"strings"

	"gitea.bee.anarckk.me/anarckk/go_util/go_file"
)

func CompressFolder(srcFolder string, dstFile string) error {
	folder, err := os.Open(srcFolder)
	if err != nil {
		return err
	}
	fileAbsPathList, err := folder.ReadDir(-1)
	if err != nil {
		return err
	}

	d, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	var files []*os.File
	for _, p := range fileAbsPathList {
		f, err := os.Open(path.Join(srcFolder, p.Name()))
		if err != nil {
			return err
		}
		files = append(files, f)
	}

	for _, file := range files {
		err := compress(file, go_file.GetFileName(srcFolder), tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// 解压 tar.gz
func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := path.Join(dest, hdr.Name)
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
