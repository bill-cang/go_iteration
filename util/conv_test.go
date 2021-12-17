/*
@Time   : 2021/10/22 9:42
@Author : ckx0709
@Remark : 转换
*/
package util

import (
	"fmt"
	"testing"
)

func TestConv2DEC(t *testing.T) {
	dec := Conv2DEC("1011")
	fmt.Print(dec)
}

func TestConv2DEC2(t *testing.T) {
	fmt.Printf("%.2f\n", 1.479-0.005)
	fmt.Printf("%.2f\n", 1.475-0.005)
	fmt.Printf("%.2f\n", 1.470-0.005)
}

