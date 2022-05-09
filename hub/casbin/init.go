/*
@Time   : 2022/3/7 14:17
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v2"
)

var GlobalEnforcer *casbin.Enforcer

func init0() {
	//mysql
	Adapter, err := gormAdapter.NewAdapter("mysql", "community:eps_community_8@tcp(192.168.6.101:3306)/community_dev", true)
	if err != nil {
		fmt.Print(err)
		return
	}

	var modelType string
	//rbac
	//modelType = `D:\WorkSpace\go_iteration\hub\casbin\config\rbac_model.conf`
	//rbac_domain
	modelType = `D:\WorkSpace\go_iteration\hub\casbin\config\rbac_domain_model.conf`

	GlobalEnforcer, err = casbin.NewEnforcer(modelType, Adapter, true)
	if err != nil {
		fmt.Print(err)
		return
	}

	//开启日志
	//GlobalEnforcer.EnableLog(true)

	//添加自定义适配器
	//GlobalEnforcer.AddFunction("customMatcher1", RegxSub)

	//从DB加载策略
	err = GlobalEnforcer.LoadPolicy()
	if err != nil {
		fmt.Println(err)
		return
	}
}
