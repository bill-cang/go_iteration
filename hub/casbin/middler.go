/*
@Time   : 2022/3/7 14:20
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

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
		fmt.Printf("domain : %s ,GetDomainsForUser: %+v\n", qr.Sub, domain)
		adapter := e.GetAdapter()
		fmt.Printf("adapter : %s ,GetAdapter: %+v\n", qr.Sub, adapter)
		forUser := e.GetPermissionsForUser(qr.Sub, roles...)
		for i := 0; i < len(forUser); i++ {
			fmt.Printf("roles : %s ,GetPermissionsForUser: %+v\n", qr.Sub, forUser[i])
		}

		//判断策略中是否存在0
		//aa, err := e.Enforce(qr.Sub, obj, act)
		aa, err := e.Enforce(qr.Sub, obj, act, true)
		aa, err = e.Enforce("admin", "community_r_1", "/api/1/visitor", "GET")
		if err != nil {
			_ = c.Error(err)
			fmt.Println("e.Enforce error")
		}
		if aa {
			fmt.Println("通过权限")
			c.Next()
		} else {
			fmt.Println("权限没有通过")
			c.Next()
		}
	}
}
