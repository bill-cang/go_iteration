/*
@Time   : 2021/11/16 12:13
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("***********************************************TOPIC: %s\n", msg.Topic())
	fmt.Printf("***********************************************MSG  : %s\n", msg.Payload())
}

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.6.120:4001").
		SetClientID(fmt.Sprintf("community-health-device:%d", 123456)).
		SetUsername("admin").
		SetPassword("Airobot@emqx202103")
/*	opts := mqtt.NewClientOptions().AddBroker("tcp://47.98.231.220:1883").
		SetClientID(fmt.Sprintf("community-health-device:%d", 123456)).
		SetUsername("adminCKX002").
		SetPassword("ZGZYZZJ-001-CKX")*/

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go subRobotCommander(c)
	go pubServer(c)

	// 断开连接
	defer c.Disconnect(250)
	time.Sleep(3 * time.Minute)
}

//
func subRobotCommander(c mqtt.Client) {
	//healthbot/to_robot/376e54a23d014312b28fa52a6d91e1bd/log
	//topic := "/healthbot/to_robot/376e54a23d014312b28fa52a6d91e1bd"
	topic := "/healthbot/to_web/67dafb6e2a15468f9a09f13c2922a5e5/visitor"
	c.Unsubscribe(topic)
	// 订阅主题
	for i := 0; i < 100; i++ {
		if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println("############", token.Error())
		}
		<-time.After(3 * time.Second)
	}

	c.Unsubscribe(topic)
}

func pubServer(c mqtt.Client) {
	topic := "/healthbot/to_service/1b3e632bc4324a9c8ac1930988f294af/376e54a23d014312b28fa52a6d91e1bd/status"
	for i := 0; i < 100; i++ {
		// 发布消息
		token := c.Publish(topic, 0, false, "online")
		token.Wait()
		<-time.After(3 * time.Second)
	}
}
