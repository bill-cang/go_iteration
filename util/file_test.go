/*
@Time   : 2021/12/15 18:05
@Author : ckx0709
@Remark :
*/
package util

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetZip(t *testing.T) {
	bytes, _ := ioutil.ReadFile(`C:\Users\ckx0709\Pictures\Snipaste_2021-11-24_10-45-30.png`)
	file, _ := ioutil.ReadFile(`C:\Users\ckx0709\Pictures\visitor.png`)
	/*	files := make([]FilesCompress, 0)
		files = append(files, FilesCompress{
			FileName: "1.png",
			Body:     bytes,
		}, FilesCompress{
			FileName: "2.png",
			Body:     file,
		})

		buf, err := Press(files)
		if err != nil {
			fmt.Printf("err =%+v", err)
			return
		}*/

	cf := NewCompressFile()
	cf.Put("1.png", bytes)
	cf.Put("2.png", file)

	buf, err := cf.Press()
	err = TouchFile("./test.zip", buf.Bytes())
	if err != nil {
		fmt.Printf("err =%+v", err)
	}

}
