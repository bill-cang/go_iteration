/*
@Time   : 2021/8/25 16:19
@Author : ckx0709
@Remark : 我虽然看出来不能运行，因为hello函数形参为sync.WaitGroup[结构体],copy传递，那么main函数定义的wg变量永远不可能
执行Done()操作，会一直处于阻塞，但我不知道汇报什么错。
Panic: fatal error: all goroutines are asleep - deadlock!
死锁。
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go hello(wg)
	wg.Wait()
}

func hello(wg sync.WaitGroup) {
	fmt.Println("hello")
	wg.Done()
}
