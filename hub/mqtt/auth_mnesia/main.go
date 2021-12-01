/*
@Time   : 2021/11/22 16:50
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("***********************************************TOPIC: %s\n", msg.Topic())
	fmt.Printf("***********************************************MSG  : %s\n", msg.Payload())
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	/*	bytes := sha256.Sum256([]byte("ZGZYZZJ-001-CKX"))
		//str := fmt.Sprintf("%s", bytes[:])
		toString := base64.StdEncoding.EncodeToString(bytes[:])
		fmt.Printf("str = %s\nstr1 =%s\n", toString, string(bytes[:]))*/

	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://47.98.231.220:1883").
		SetClientID("community-health-device:123456").
		SetUsername("adminCKX001").
		SetPassword("ZGZYZZJ-001-CKX")

	//使用证书登录
	//opts.SetTLSConfig(tlsConfig)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("*******conn failed . err =%+v\n", token.Error())
	}

	fmt.Printf("#####conn = %t", c.IsConnected())

	pubServer(c)

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
