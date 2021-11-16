/*
@Time   : 2021/10/20 11:26
@Author : ckx0709
@Remark :
*/
package main

import (
	"encoding/json"
	"go_iteration/util"
	"log"
	"os/user"

	"github.com/zcalusic/sysinfo"
)

func main() {
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal("requires superuser privilege")
	}

	var fileBody []byte
	defer func() {
		util.CreateFileWithSome("./ghw.txt", fileBody)
	}()

	var si sysinfo.SysInfo
	si.GetSysInfo()
	fileBody, err = json.MarshalIndent(&si, "", "  ")
	if err != nil {
		fileBody = []byte(err.Error())
	}
}
