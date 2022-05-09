/*
@Time   : 2022/3/7 14:21
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"github.com/casbin/casbin/v2/util"
	"testing"
)

//RBAC DOMAIN
//Enforce(角色RuleID[v0],域[v1],资源Router[v2],操作Action[v3])GET
func TestMiddleCasBin2(t *testing.T) {

	//domains, _ := GlobalEnforcer.GetDomainsForUser("admin")

	enforce, err := GlobalEnforcer.Enforce("admin", "visitor", "/api/234325435234/visitor", "PUT")
	if err != nil {
		fmt.Print(err)
		return
	}

	if enforce {
		fmt.Printf("Success :%t\n", enforce)
	} else {
		fmt.Printf("Falied :%t\n", enforce)
	}
}

//正则匹配
func TestMiddleCasBin3(t *testing.T) {
	match, err := util.RegexMatchFunc("visitor", "(visitor|commander)")
	if err != nil {
		return
	}
	fmt.Printf("正则匹配：%t\n", match)

	match, err = util.RegexMatchFunc("/api/123/visitor", "/api/[\\w]{1,32}/visitor")
	if err != nil {
		return
	}
	fmt.Printf("正则匹配：%t\n", match)
}

func TestMiddleCasBin4(t *testing.T) {

	//(visitor|station)
	domains, err := GlobalEnforcer.GetDomainsForUser("admin")
	if err != nil {
		return
	}
	fmt.Printf("GetDomainsForUser :%+v\n", domains)

	roles, err := GlobalEnforcer.GetRolesForUser("admin")
	if err != nil {
		return
	}
	fmt.Printf("GetRolesForUser :%+v\n", roles)

	Permissions := GlobalEnforcer.GetPermissionsForUser("admin")
	fmt.Printf("GetPermissionsForUser :%+v\n", Permissions)

	manager := GlobalEnforcer.GetRoleManager()
	mRoles, _ := manager.GetRoles("admin")
	fmt.Printf("GetRoles:%+v\n", mRoles)
}
