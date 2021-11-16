/*
@Time   : 2021/10/18 17:23
@Author : ckx0709
@Remark :
*/
package main

import (
	"go_iteration/lib"
	"os"
)

func main() {
	params := lib.GetAppDataPath()
	file, _ := os.Create("./env.txt")
	file.WriteString(params)
}
