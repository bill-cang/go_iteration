/*
@Time   : 2021/12/15 15:38
@Author : ckx0709
@Remark :
*/
package client

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

var Cli2 mqtt.Client

func init() {
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.6.120:4001").
		SetClientID("ckx-test-client").
		SetUsername("admin").
		SetPassword("Airobot@emqx202103")

	//使用证书登录
	//opts.SetTLSConfig(tlsConfig)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	//opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	Cli2 = mqtt.NewClient(opts)
	if token := Cli2.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("*******conn failed . err =%+v\n", token.Error())
	}
}
