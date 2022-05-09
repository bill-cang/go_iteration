/*
@Time   : 2022/4/24 17:33
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"regexp"
)

type IA interface {
	Eat()
	Song()
}

type IB interface {
	IA
}

type A struct {
}

func (a *A) Song() {

}

func (a *A) Eat() {

}

type B struct {
	A
}

func main() {
	//strs := make([]string, 0)

	compile := regexp.MustCompile("column:([\\w]+);?")

	findString := compile.FindStringSubmatch("column:master;comment:负责人")
	fmt.Println(findString[1])

}
