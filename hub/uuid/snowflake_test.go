/*
@Time   : 2021/12/17 10:32
@Author : ckx0709
@Remark :
*/
package uuid

import (
	"fmt"
	"testing"
)

func TestGetSnowflakeUUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d\n", GetSnowflakeUUID())
	}
}

func BenchmarkGetSnowflakeUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetSnowflakeUUID()
	}
}
