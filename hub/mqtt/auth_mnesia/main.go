/*
@Time   : 2021/12/1 10:57
@Author : ckx0709
@Remark :
*/
package main

import (
	"context"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go_iteration/hub/mqtt/auth_mnesia/client"
	"go_iteration/hub/mqtt/auth_mnesia/contest"
	"google.golang.org/appengine/log"
	"os"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for i := 0; i < 1; i++ {
			msg := fmt.Sprintf("online[%d]", i)
			token := client.MQTTClient.Publish(contest.TopicTest, 2, true, msg)
			if token.Wait() && token.Error() != nil {
				log.Errorf(context.TODO(), "Subscribe err :%+v", token.Error())
			}
			fmt.Printf("[Publish] send %s\n", msg)
			time.Sleep(3 * time.Second)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			token := client.MQTTClient.Subscribe(contest.TopicTest, 2, nil)
			if token.Wait() && token.Error() != nil {
				log.Errorf(context.TODO(), "Subscribe err :%+v", token.Error())
			}
			time.Sleep(5 * time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
}

func toFile(client mqtt.Client, msg mqtt.Message) {
	file, _ := os.OpenFile("qos2.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	defer file.Close()
	temp := fmt.Sprintf("messageID = %d,payload = %s, topic = %s\n", msg.MessageID(), msg.Payload(), msg.Topic())
	file.WriteString(temp)
}
