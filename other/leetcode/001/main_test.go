/*
@Time   : 2022/4/15 16:39
@Author : ckx0709
@Remark :
*/
package main

import "testing"

func BenchmarkFindIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindIndex(numbers, target)
	}
}

func BenchmarkFindIndex2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindIndex2(numbers, target)
	}
}
