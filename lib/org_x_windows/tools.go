// Package org_x_windows /*
package org_x_windows

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/registry"
	"io"
	"io/ioutil"
	"net"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GetWindowsStopTimeout 获取Windows上一次终止服务时间
func GetWindowsStopTimeout() time.Duration {
	defaultTimeout := time.Millisecond * 20000
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control`,
		registry.READ)
	if err != nil {
		return defaultTimeout
	}
	sv, _, err := key.GetStringValue("WaitToKillServiceTimeout")
	if err != nil {
		return defaultTimeout
	}
	v, err := strconv.Atoi(sv)
	if err != nil {
		return defaultTimeout
	}
	return time.Millisecond * time.Duration(v)
}

// StartWinQQ 唤醒Windows qq进程
func StartWinQQ() {
	//读取注册表
	key, err := registry.OpenKey(registry.CLASSES_ROOT, `Tencent\shell\open\command`,
		registry.READ)
	if err != nil {
		return
	}
	//获取键的默认值
	sv, _, err := key.GetStringValue("")
	//
	compile := regexp.MustCompile(`"([^"]+)"+?`)
	if err != nil {
		fmt.Println(err)
		return
	}
	findString := compile.FindStringSubmatch(sv)
	//正则将qq bin目录启动文件替换成qq启动
	mustCompile := regexp.MustCompile(`([\w])+.exe`)

	//组装启动命令
	//func1
	command := exec.Cmd{
		Path: mustCompile.ReplaceAllString(findString[1], "QQ.exe"),
	}

	//func2
	/*	allString := mustCompile.ReplaceAllString(findString[1], "QQ.exe")
		command := exec.Command(allString)*/
	//启动QQ
	err = command.Start()
	if err != nil {
		println(err)
		return
	}
}

func listenPort(port int) {
	listen, err := net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Error(err)
		return
	}

	for true {
		conn, err := listen.Accept()
		if err != nil {
			log.Error(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		log.Fatalf("[handleConnection] %s", data)
		fmt.Fprintf(conn, "OK\n")
	}
	conn.Close()
}

// WriteStartUp 写入开机自启动
func WriteStartUp(souExe string) {
	bytes, err := ioutil.ReadFile(souExe)
	if err != nil {
		log.Error(err)
		return
	}
	startUpDir := getUserStartUpDir()
	if startUpDir == "" {
		return
	}
	err = ioutil.WriteFile(startUpDir, bytes, 438)
	if err != nil {
		log.Error(err)
		return
	}
}

//获取用户自启目录
func getUserStartUpDir() string {
	current, err := user.Current()
	if err != nil {
		log.Error(err)
		return ""
	}

	homeDir := strings.Replace(current.HomeDir, `\`, "/", -1)
	builder := strings.Builder{}
	builder.WriteString(homeDir)
	builder.WriteString(fmt.Sprintf("/%s/%s.exe", startABS, asName))
	return builder.String()
}
