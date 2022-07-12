/*
@Time   : 2022/7/7 9:32
@Author : ckx0709
@Remark :
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())
	valCtx := context.WithValue(ctx, "position", "gopher")
	wg.Add(2)
	go func() {
		defer wg.Done()
		worker(valCtx, "打工人1")
	}()
	go func() {
		defer wg.Done()
		worker(valCtx, "打工人2")
	}()
	time.Sleep(3 * time.Second) //工作3秒
	stop()                      //3秒后发出停止指令
	wg.Wait()
}

func worker(valCtx context.Context, name string) {
	for {
		select {
		case <-valCtx.Done():
			fmt.Println("下班咯~~~")
			return
		default:
			position := valCtx.Value("position")
			fmt.Println(name, position, "认真摸鱼中，请勿打扰...")
		}
		time.Sleep(1 * time.Second)
	}
}
