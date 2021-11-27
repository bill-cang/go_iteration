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

func main() {
/*	bytes := md5.Sum([]byte("ZGZYZZJ-001-CKX"))
	fmt.Printf("sum =%s\n", bytes)*/

	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://47.98.231.220:1883").
		SetClientID("community-health-device:123456").
		SetPassword("ZGZYZZJ-001-CKX").
		SetUsername("admin").
		SetPassword("ZGZYZZJ-001-CKX")

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("*******conn failed . err =%+v\n", token.Error())
	}

	fmt.Printf("#####conn = %t", c.IsConnected())

}
