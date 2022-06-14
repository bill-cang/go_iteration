/*
@Time   : 2022/6/8 9:32
@Author : ckx0709
@Remark :
sync.Cond 实现一个变量条件，控制唤醒一组goroutines得场景。
实现发布订阅
*/
package main

import (
	"log"
	"sync"
	"time"
)

type ExChange struct {
	Done bool
	Cond *sync.Cond
}

func NewExchange() ExChange {
	return ExChange{
		Done: false,
		Cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (e *ExChange) Publish(do func()) {
	e.Cond.L.Lock()
	do()
	e.Cond.L.Unlock()
	log.Println("Publish done")
	e.Done = true
	//唤醒全部等待的goroutine
	e.Cond.Broadcast()
}

func (e *ExChange) Subscribe(do func()) {
	e.Cond.L.Lock()
	for !e.Done {
		e.Cond.Wait()
	}
	do()
	e.Cond.L.Unlock()
}

func main() {
	//公共资源
	mmp := map[int]string{0: "A"}

	//修改公共资源
	pubFun := func() {
		mmp[0] = "a"
	}

	subFun := func() {
		for i, s := range mmp {
			log.Printf("%d :%s", i, s)
		}
	}

	TXBan := NewExchange()
	for i := 0; i < 10; i++ {
		go TXBan.Subscribe(subFun)
	}

	TXBan.Publish(pubFun)
	time.Sleep(3 * time.Second)
}
