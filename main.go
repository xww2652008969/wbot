package main

import (
	"fmt"
	"github.com/xww2652008969/wbot/client"
)

func main() {
	config := client.Clientconfig{
		Wsurl:      "ws://192.168.10.209",
		Wspost:     "3001",
		Wsheader:   nil,
		Clienthttp: "http://192.168.10.209:4000",
	}
	c, err := client.New(config)
	if err != nil {
		panic(err)
	}
	c.RegisterPush(te())
	c.RegisterGroupHandle(te2())
	c.RegisterPrivateHandle(te3())
	c.RegisterMessageSenthandle(te4())
	c.RegisterInterceptor(Chack())
	c.Run()
}

func Chack() client.Interceptorfunc {
	return func(c client.Client, message client.Message) bool {
		fmt.Println("触发了拦截,但是放行")
		return true
	}
}

func te() client.Push {
	return func(client client.Client) {
		client.AddText("主动执行").AddImage("1").AddText("1")
		client.SendPrivateMsg(1271701079)
	}
}
func te2() client.Event {
	return func(client client.Client, message client.Message) {
		fmt.Println("触发了消息")
	}
}
func te3() client.Event {
	return func(client client.Client, message client.Message) {
		fmt.Println(message)
		var a = client.AddText("这是私聊事件")
		fmt.Println(a)
		client.SendPrivateMsg(1271701079)
	}
}
func te4() client.Event {
	return func(client client.Client, message client.Message) {
		fmt.Println(message)
		fmt.Println("触发主动发送")
		client.DeleMsg(message.MessageId)
	}
}
