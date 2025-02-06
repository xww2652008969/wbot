package main

import (
	"github.com/xww2652008969/wbot/MessageType"
	"github.com/xww2652008969/wbot/client"
	"github.com/xww2652008969/wbot/client/Message"
)

func main() {
	config := client.Clientconfig{
		Wsurl:      "ws://192.168.10.209",
		Wspost:     "3001",
		Wsheader:   nil,
		Clienthttp: "http://192.168.10.209:4000",
	}
	c, err := client.Create(config)
	if err != nil {
		panic(err)
	}
	c.Register(MessageType.Group, te1())
	c.Register(MessageType.Private, te2())
	c.Run()
	select {}

}
func te1() Message.Event {
	return func(message Message.Message) {
		message.AddText("测试群私聊的")
		message.SendGPrivateMsg(3555202423)
	}
}
func te2() Message.Event {
	return func(message Message.Message) {
	}
}
