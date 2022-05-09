/*
@Time   : 2022/4/12 17:32
@Author : ckx0709
@Remark :
*/
package main

import "fmt"

type rangeType interface {
	[]byte | map[int]byte | string
}

func rangeIt[R rangeType](r R) {
	for i, v := range r {
		fmt.Println(i, v)
	}
}

func main() {
	rangeIt(map[int]byte{1: 1, 2: 2, 3: 3})
}
