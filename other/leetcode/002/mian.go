/*
@Time   : 2022/4/24 17:33
@Author : ckx0709
@Remark :
*/
package main

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
