/*
@Time   : 2021/12/28 11:58
@Author : ckx0709
@Remark :
*/
package main

import (
	"context"
	"fmt"
)

// 返回⽣成⾃然数序列的管道: 2, 3, 4, ...
func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func main() {
	// 通过 Context 控制后台Goroutine状态
	ctx, cancel := context.WithCancel(context.Background())
	ch := GenerateNatural(ctx)
	// ⾃然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch
		// 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ctx, ch, prime)
		// 基于新素数构造的过滤器
	}
	cancel()
}
