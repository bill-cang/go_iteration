/*
@Time   : 2021/10/20 10:42
@Author : ckx0709
@Remark : 获取计算机硬件信息
*/
package main

import (
	"encoding/json"
	"github.com/jaypipes/ghw"
	"go_iteration/util"
)

func main() {
	cpu, err := ghw.CPU(ghw.WithChroot("/host"))

	var fileBody []byte
	defer func() {
		util.CreateFileWithSome("./ghw.txt", fileBody)
	}()

	if err != nil {
		fileBody = []byte(err.Error())
	}

	fileBody, err = json.Marshal(cpu)
	if err != nil {
		fileBody = []byte(err.Error())
	}
}
