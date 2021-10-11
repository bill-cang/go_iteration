/*
@Time   : 2021/10/11 14:32
@Author : ckx0709
@Remark :
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
	a, err := gormadapter.NewAdapter("mysql", "ckx:passwordbyCKX@tcp(47.98.231.220:3306)/registrar",true)
	if err != nil {
		fmt.Print(err)
		return
	}
	e, _ := casbin.NewEnforcer("D:\\WorkSpace\\go_iteration\\hub\\casbin\\model.conf", a)

	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")

}
