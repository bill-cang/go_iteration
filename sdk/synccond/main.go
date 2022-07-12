/*
@Time   : 2022/6/8 9:32
@Author : ckx0709
@Remark :
sync.Cond 实现一个变量条件，控制唤醒一组goroutines得场景。
特性：
sync.Cond一旦创建使用 不允许被拷贝，由noCopy和copyChecker来限制保护。
Wait()操作先是递增notifyList.wait属性 然后将goroutine封装进sudog，将notifyList.wait赋值给sudog.ticket,然后将sudog插入notifyList链表中
Singal()实际是按照notifyList.notify跟notifyList链表中节点的ticket匹配 来确定唤醒的goroutine，因为notifyList.notify和notifyList.wait都是原子递增的，故而有了FIFO的语义
Broadcast()相对简单 就是唤醒全部等待的goroutine

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
