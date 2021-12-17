/*
@Time   : 2021/12/15 18:04
@Author : ckx0709
@Remark : 文件压缩
*/
package util

import (
	"archive/zip"
	"bytes"
)

//压缩文件实体map[文件名]文件字节流
type FilesCompress struct {
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

//输出压缩buffer
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

//输出压缩文件
func (c FilesCompress) Touch(filePath string) (err error) {
	buf, err := c.Press()
	if err != nil {
		return err
	}
	return FileTouch(filePath, buf.Bytes())
}
