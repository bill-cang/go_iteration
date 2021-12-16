/*
@Time   : 2021/12/15 18:04
@Author : ckx0709
@Remark :
*/
package util

import (
	"archive/zip"
	"bytes"
)

//压缩文件
type FilesCompress struct {
	/*	FileName string //文件名
		Body     []byte //文件内容*/
	Files map[string][]byte
}

//初始化
func NewCompressFile() FilesCompress {
	return FilesCompress{
		Files: make(map[string][]byte),
	}
}

//添加文件
func (c *FilesCompress) Put(fileName string, fileBody []byte) {
	c.Files[fileName] = fileBody
}

//输出压缩buf
func (c *FilesCompress) Press() (buf *bytes.Buffer, err error) {
	buf = new(bytes.Buffer)
	w := zip.NewWriter(buf)
	defer w.Close()
	for fileName, fileBody := range c.Files {
		writer, err := w.Create(fileName)
		if err != nil {
			return nil, err
		}
		_, err = writer.Write(fileBody)
		if err != nil {
			return nil, err
		}
	}
	return
}
