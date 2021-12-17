/*
@Time   : 2021/12/1 10:24
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

var (
	MQTTClient mqtt.Client

	messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s ;from topic: %s\n", msg.Payload(), msg.Topic())
	}

	connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Printf("Connected success.\n")
	}

	connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}
)

func init() {
	/*	bytes := sha256.Sum256([]byte("ZGZYZZJ-001-CKX"))
		//str := fmt.Sprintf("%s", bytes[:])
		toString := base64.StdEncoding.EncodeToString(bytes[:])
		fmt.Printf("str = %s\nstr1 =%s\n", toString, string(bytes[:]))*/

	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://47.98.231.220:1883").
		SetClientID("SN:1").
		SetUsername("adminCKX001").
		SetPassword("ZGZYZZJ-001-CKX")

	//使用证书登录
	//opts.SetTLSConfig(tlsConfig)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	//opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	MQTTClient = mqtt.NewClient(opts)
	if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("*******conn failed . err =%+v\n", token.Error())
	}
}
