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
