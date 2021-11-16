/*
@Time   : 2021/10/20 10:45
@Author : ckx0709
@Remark :
*/
package util

import "os"

var (
	defFlag                  = os.O_CREATE | os.O_TRUNC | os.O_RDWR
	defFileModel os.FileMode = 0777
)

/*创建文件
params[0] flag = 0 maybe default
params[0] FileMode = nil is 0777
*/
func CreateFileWithSome(fileName string, bytes []byte, params ...interface{}) error {
	file, err := os.OpenFile(fileName, defFlag, defFileModel)
	if err != nil {
		return err
	}
	_, err = file.Write(bytes)
	return err
}
