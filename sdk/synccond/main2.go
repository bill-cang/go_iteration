/*
@Time   : 2022/7/11 11:40
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var cond *sync.Cond

func init() {

	cond = sync.NewCond(&sync.Mutex{})
}

func test(x int) {
	cond.L.Lock() // 获取锁
	cond.Wait()   // 等待通知 暂时阻塞
	fmt.Println(x)
	time.Sleep(time.Second * 1)
	cond.L.Unlock()
}

func CondTest1() {
	for i := 0; i < 40; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 3)
	fmt.Println("broadcast")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Broadcast() //3秒之后 下发广播给所有等待的goroutine
}
