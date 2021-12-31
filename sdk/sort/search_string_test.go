/*
@Time   : 2021/12/29 15:44
@Author : ckx0709
@Remark :
*/
package sort

import (
	"fmt"
	"net/url"
	"testing"
)

func TestSearchStringFromSplice(t *testing.T) {
	ss := []string{"a", "v", "s", "f", "r"}
	ind := SearchStringFromSplice(ss, "z")
	fmt.Println(ind)
}

/*
circular [ˈsɜːrkjələr]	圆形的，环形的		记忆技巧：circ 圆，环 + ular 有…形状或性质的 → 圆形的
oval 	椭圆形的		the oval office 椭圆办公室->隐身指美国总统
ark		弧形
segment		弓形		orange segments 橘子瓣
*/

func TestSearchStringFromSplice2(t *testing.T) {
	escape := url.QueryEscape("2021-12-24T17:18:01+08:00")
	fmt.Println(escape)
}

