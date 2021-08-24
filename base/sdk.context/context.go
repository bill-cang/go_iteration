/*
@Time   : 2021/8/21 17:04
@Author : ckx0709
@Remark :
*/
package main

import (
	"context"
	"fmt"
	"time"
)

// parentGoroutine 父G
func parentGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号，父ctx err=", ctx.Err())
			return
		default:
			fmt.Println("父ctx监控中")
			time.Sleep(1 * time.Second)
		}
	}
}

// childrenGoroutine 子G
func childrenGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号，子ctx err=", ctx.Err())
			return
		default:
			fmt.Println("子ctx监控中")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 创建一个根节点
	ctx := context.Background()

	// 可退出的Context
	parentCtx, cancelFunc := context.WithCancel(ctx)
	// 监控parentCtx 的 done 通知
	go parentGoroutine(parentCtx)

	timerChildrenCtx, _ := context.WithCancel(parentCtx) // 收到父G信号，退出
	// 监控childrenCtx 的 done 通知
	go childrenGoroutine(timerChildrenCtx)

	//time.Sleep(3 * time.Second)
	cancelFunc() // 通知parentCtx，退出

	time.Sleep(6 * time.Second)
}
