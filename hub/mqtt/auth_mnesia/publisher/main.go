/*
@Time   : 2021/12/1 10:54
@Author : ckx0709
@Remark :
*/
package main

import (
	"context"
	"fmt"
	"go_iteration/hub/mqtt/auth_mnesia/client"
	"go_iteration/hub/mqtt/auth_mnesia/contest"
	"google.golang.org/appengine/log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	subscriber()

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			msg := fmt.Sprintf("online[%d]", i)
			token := client.MQTTClient.Publish(contest.TopicTest, 2, true, msg)
			if token.Wait() && token.Error() != nil {
				log.Errorf(context.TODO(), "Subscribe err :%+v", token.Error())
				break
			}
			fmt.Printf("[Publish] send %s\n", msg)
			time.Sleep(3 * time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
}

func subscriber() {
	go func() {
		topic := "/healthbot/to_service/12/+/status"
		token := client.Cli2.Subscribe(topic, 2, nil)
		if token.Wait() && token.Error() != nil {
			log.Errorf(context.TODO(), "Subscribe err :%+v", token.Error())
		}
	}()
}
