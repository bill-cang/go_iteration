/*
@Time   : 2021/8/21 18:27
@Author : ckx0709
@Remark :
reflect.Value.CanAddr():判断是否可寻址
reflect.Type.Elem():其反射对象上调用获取它指向的元素的reflect.Value
*/
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := &User{Name: "dj", Age: 18}
	t := reflect.TypeOf(u).Elem()
	fmt.Println(t)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Tag)
	}
}
