/*
@Time   : 2021/12/1 10:55
@Author : ckx0709
@Remark :
*/
package main

import (
	"context"
	"go_iteration/hub/mqtt/auth_mnesia/client"
	"go_iteration/hub/mqtt/auth_mnesia/contest"
	"google.golang.org/appengine/log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		token := client.Cli2.Subscribe(contest.TopicTest, 2, nil)
		if token.Wait() && token.Error() != nil {
			log.Errorf(context.TODO(), "Subscribe err :%+v", token.Error())
		}

		wg.Done()
	}()

	time.Sleep(300 * time.Second)
	wg.Wait()
}
