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
	_ "github.com/go-sql-driver/mysql"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {

	//mysql
	Enforcer, err := gormadapter.NewAdapter("mysql", "ckx:pwdByCkx*@tcp(47.98.231.220:3306)/registrar", true)
	if err != nil {
		fmt.Print(err)
		return
	}

	e, _ := casbin.NewEnforcer("D:\\WorkSpace\\go_iteration\\hub\\casbin\\model.conf", Enforcer)

	checkRoot(e)
	checkRole(e)
	checkUser(e)
}

//校验-超级账户无碍通过
func checkRoot(e *casbin.Enforcer) {
	enforce, err := e.Enforce("root", "object", "read")
	fmt.Printf("enforce =%t, err =%+v\n", enforce, err)
}

//校验-角色匹配通过
func checkRole(e *casbin.Enforcer) {
	enforce2, err2 := e.Enforce("engineer", "object", "read")
	fmt.Printf("enforce =%t, err2 =%+v\n", enforce2, err2)
}

//校验-用户匹配通过
func checkUser(e *casbin.Enforcer) {
	enforce, err := e.Enforce("ckx", "object", "read")
	fmt.Printf("enforce =%t, err =%+v\n", enforce, err)
}
