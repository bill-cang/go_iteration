/*
@Time   : 2022/4/8 9:35
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	//gormAdapter "github.com/casbin/gorm-adapter/v2"
	"go_iteration/hub/casbin/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	//community:eps_community_8@tcp(192.168.6.101:3306)/community_dev
	dsn := "community:eps_community_8@tcp(192.168.6.101:3306)/community_dev"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		return
	}

	db.Where("")

	// Initialize adapter Gorm adapter and use it in adapter Casbin enforcer:
	// The adapter will use an existing gorm.DB instnace.
	casRule := entity.CasRule{}
	adapter, err := gormAdapter.NewAdapterByDBWithCustomTable(db, &casRule, casRule.TableName())
	if err != nil {
		fmt.Print(err)
		return
	}

	var (
		modelType string
		//err       error
	)
	//rbac
	//modelType = `D:\WorkSpace\go_iteration\hub\casbin\config\rbac_model.conf`
	//rbac_domain
	modelType = `D:\WorkSpace\go_iteration\hub\casbin\config\rbac_domain_model.conf`
	GlobalEnforcer, err = casbin.NewEnforcer(modelType, adapter, true)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Load the policy from DB.
	err = GlobalEnforcer.LoadPolicy()
	if err != nil {
		fmt.Print(err)
		return
	}

}
