/*
@Time   : 2021/10/11 14:32
@Author : ckx0709
@Remark :
casbin API: https://casbin.org/docs/zh-CN
*/
package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//mysql
	Adapter, err := gormadapter.NewAdapter("mysql", "community:eps_community_8@tcp(192.168.6.101:3306)/community_dev", true)
	if err != nil {
		fmt.Print(err)
		return
	}

	var modelType string
	//rbac
	//modelType = `D:\WorkSpace\go_iteration\hub\casbin\config\rbac_model.conf`
	//rbac_domain
	modelType = `D:\WorkSpace\go_iteration\hub\casbin\config\rbac_domain_model.conf`

	enforcer, err := casbin.NewEnforcer(modelType, Adapter)
	if err != nil {
		fmt.Print(err)
		return
	}

	//开启日志
	enforcer.EnableLog(true)

	//添加自定义适配器
	//enforcer.AddFunction("customMatcher1", RegxSub)

	//从DB加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		fmt.Println(err)
		return
	}

	/*	adapter := enforcer.GetAdapter()
		fmt.Println()*/

	//获取router路由对象
	r := gin.New()
	//使用自定义拦截器中间件
	r.Use(MiddleCasBin(enforcer))
	//注册路由
	registerRouter(r)

	err = r.Run(":8341")
	//参数为空 默认监听8080端口

}

//拦截器
func MiddleCasBin(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {

		//获取请求的URI
		obj := c.Request.URL.Path
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		type Query struct {
			Sub string
		}
		qr := new(Query)
		err := c.Bind(&qr)
		if err != nil {
			c.Abort()
		}

		roles, _ := e.GetRolesForUser(qr.Sub)
		fmt.Printf("roles : %s ,GetRolesForUser: %s\n", qr.Sub, roles)
		domain, _ := e.GetDomainsForUser(qr.Sub)
		fmt.Printf("roles : %s ,GetDomainsForUser: %s\n", qr.Sub, domain)
		adapter := e.GetAdapter()
		fmt.Printf("roles : %s ,GetAdapter: %+v\n", qr.Sub, adapter)
		forUser := e.GetPermissionsForUser(qr.Sub, roles...)
		for i := 0; i < len(forUser); i++ {
			fmt.Printf("roles : %s ,GetPermissionsForUser: %+v\n", qr.Sub, forUser[i])
		}

		//判断策略中是否存在
		aa, err := e.Enforce(qr.Sub, obj, act)
		if err != nil {
			fmt.Println("e.Enforce error")
		}
		if aa {
			fmt.Println("通过权限")
			c.Next()
		} else {
			fmt.Println("权限没有通过")
			c.Abort()
		}
	}
}
