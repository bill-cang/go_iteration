/*
@Time   : 2022/6/24 11:23
@Author : ckx0709
@Remark :
*/
package org_x_windows

import (
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func TestGetWindowsStopTimeout(t *testing.T) {
	open, _ := os.Open(`F:\GoPath\bin\sloth.exe`)
	stat, _ := open.Stat()
	log.Fatalf("chomd :%d", stat.Mode())
}

func TestGetWindowsStopTimeout2(t *testing.T) {
	StartWinQQ()
}

func TestGetWindowsStopTimeout3(t *testing.T) {
	dll, err := syscall.LoadDLL(`D:\Program Files (x86)\Tencent\QQ\Bin\AddrSearch.dll`)
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	fmt.Println(dll)

}

func TestGetWindowsStopTimeout4(t *testing.T) {
	listenPort(8000)
}

func TestGetWindowsStopTimeout5(t *testing.T) {
	//C:\Users\ckx0709\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup
	WriteStartUp(`F:\GoPath\bin\swag.exe`)
}

func TestStartWinQQ(t *testing.T) {

	command := exec.Command("ls")
	//command.Args = []string{"all"}
	/*	err := command.Start()
		if err != nil {
			log.Error(err)
			return
		}*/
	command.Env = append(os.Environ(), "GOOS=linux")
	bytes, err := command.Output()
	if err != nil {
		log.Error(err)
		return
	}
	log.Printf("output :%s", bytes)

	buf, _ := command.CombinedOutput()
	log.Println(string(buf))

}

func TestStartWinQQ2(t *testing.T) {
	pwd := "RVBTLWNreA=="
	bytes, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		panic(err)
	}
	log.Printf("pwd =%s", bytes)
}

func TestWriteStartUp(t *testing.T) {
}
