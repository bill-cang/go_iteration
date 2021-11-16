/*
@Time   : 2021/10/18 14:13
@Author : ckx0709
@Remark : 获取系统用户环境配置
*/
package lib

import (
	"os"
	"runtime"
	"strings"
)

var (
	appDataPath string
)

//获取环境变量
func GetEnvParams() []string {
	return os.Environ()
}

//获取指定环境变量值
func GetEnvParamByKey(key string) string {
	envParams := GetEnvParams()
	for _, s := range envParams {
		if index := strings.Index(s, key); index == 0 {
			return strings.Split(s, "=")[1]
		}
	}
	return ""
}

//获取appdata dir
func GetAppDataPath() string {
	if appDataPath == "" {
		var key string
		switch runtime.GOOS {
		case "linux":
			return "/usr/local/etc"
		case "windows":
			key = "APPDATA"
		}
		appDataPath = GetEnvParamByKey(key)
	}
	return appDataPath
}
