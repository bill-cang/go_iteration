/*
@Time   : 2021/12/17 10:22
@Author : ckx0709
@Remark :
*/
package uuid

import (
	"fmt"
	"testing"
)

func TestUUIDOrderly(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s\n", GetOrderlyUUID())
	}
}

func BenchmarkUUIDOrderly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetOrderlyUUID()
	}
}
