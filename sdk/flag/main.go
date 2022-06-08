/*
@Time   : 2022/6/2 10:42
@Author : ckx0709
@Remark :
https://juejin.cn/post/7098245145744965663
go run main.go -host=localhost -user=test -password=123456 -db_name=test -port=3307
另外，运行程序时，在后面跟上-h或--help来查看命令的参数选项
*/
package main

import (
	"flag"
	"fmt"
)

var (
	host     string
	dbName   string
	port     int
	user     string
	password string
)

func main() {

	//接收参数的变量\参数名称\默认值\参数说明\
	flag.StringVar(&host, "host", "", "数据库地址")
	flag.StringVar(&dbName, "db_name", "", "数据库名称")
	flag.StringVar(&user, "user", "", "数据库用户")
	flag.StringVar(&password, "password", "", "数据库密码")
	//设置默认值
	flag.IntVar(&port, "port", 3306, "数据库端口")

	flag.Parse()

	fmt.Printf("数据库地址:%s\n", host)
	fmt.Printf("数据库名称:%s\n", dbName)
	fmt.Printf("数据库用户:%s\n", user)
	fmt.Printf("数据库密码:%s\n", password)
	fmt.Printf("数据库端口:%d\n", port)

}
