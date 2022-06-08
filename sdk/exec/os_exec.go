/*
@Time   : 2022/6/2 10:01
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")

	e := cmd.Run()
	CheckError(e)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
