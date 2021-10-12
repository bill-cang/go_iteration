/*
@Time   : 2021/10/11 14:32
@Author : ckx0709
@Remark :
casbin API: https://casbin.org/docs/zh-CN
*/
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"

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
	/*	Enforcer, err := gormadapter.NewAdapter("mysql", "ckx:@tcp(47.98.231.220:3306)/registrar", true)
		if err != nil {
			fmt.Print(err)
			return
		}*/

	//gorm
	dsn := "ckx:@tcp(47.98.231.220:3306)/registrar"
	dbconn, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect DB error")
		panic(err)
	}
	// 将数据库连接同步给插件， 插件用来操作数据库
	adt, err := gormadapter.NewAdapterByDB(dbconn)
	if err != nil {
		fmt.Println("connect DB error")
		panic(err)
	}
	// 这里也可以使用原生字符串方式
	//
	Enforcer, _ := casbin.NewEnforcer("./model.conf", adt)

	e, _ := casbin.NewEnforcer("D:\\WorkSpace\\go_iteration\\hub\\casbin\\model.conf", Enforcer)

	enforce, err := e.Enforce("dajun", "data1", "read")
	fmt.Printf("enforce =%t, err =%+v\n", enforce, err)
	enforce2, err2 := e.Enforce("dajun", "data2", "read")
	fmt.Printf("enforce2 =%t, err2 =%+v\n", enforce2, err2)

	/*	check(e, "lizi", "data1", "write")
		check(e, "lizi", "data2", "read")
		check(e, "dajun", "data1", "write")
		check(e, "dajun", "data2", "read")*/

}
