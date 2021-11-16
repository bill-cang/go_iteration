/*
@Time   : 2021/10/21 9:57
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

func main() {
	out := 0
	fc := func() {
		out++
		fmt.Println(out)
	}

	c := cron.New(cron.WithSeconds())
	taskID, err := c.AddFunc("*/3 * * * * ?", fc)
	fmt.Printf("taskID = %+v, err= %+v\n", taskID, err)

	c.Start()
	time.Sleep(15 * time.Second)
}

func main1() {
	can := make(chan int)

	out := 0
	fc := func() {
		out++
		fmt.Println(out)
		if out == 5 {
			can <- out
		}
	}

	c := cron.New(cron.WithSeconds())
	taskID, err := c.AddFunc("*/3 * * * * ?", fc)
	fmt.Printf("taskID = %+v, err= %+v\n", taskID, err)

	//缺失c.Start()：fatal error: all goroutines are asleep - deadlock!
	c.Start()
	<-can
}

func main2() {
	wait := &sync.WaitGroup{}
	wait.Add(5)

	out := 0
	fc := func() {
		out++
		fmt.Println(out)
		wait.Done()
	}

	c := cron.New(cron.WithSeconds())
	taskID, err := c.AddFunc("*/3 * * * * ?", fc)
	fmt.Printf("taskID = %+v, err= %+v\n", taskID, err)

	//缺失c.Start()：fatal error: all goroutines are asleep - deadlock!
	c.Start()
	wait.Wait()
}
