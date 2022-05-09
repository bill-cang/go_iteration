/*
@Time   : 2022/4/12 17:46
@Author : ckx0709
@Remark :
*/
package main

type Duck interface {
	Song(string)
}

type Dog interface {
	Song(string)
}

type Animal interface {
	Duck
	Dog
}

type DaHuang struct {
}

//显式实现
var _ Animal = (*DaHuang)(nil)

func (s *DaHuang) Song(string2 string) {

}
