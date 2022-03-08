/*
@Time   : 2021/10/11 14:32
@Author : ckx0709
@Remark :
casbin API: https://casbin.org/docs/zh-CN
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*var GlobalEnforcer *casbin.Enforcer*/

func main() {
	//获取router路由对象
	r := gin.New()
	//使用自定义拦截器中间件
	r.Use(MiddleCasBin(GlobalEnforcer))
	//注册路由
	registerRouter(r)

	routes := r.Routes()
	fmt.Printf("routes :%+v", routes)

	_ = r.Run(":8341")
	//参数为空 默认监听8080端口 有栈携程和无栈携程

}
