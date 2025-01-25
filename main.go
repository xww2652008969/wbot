package main

import (
	"fmt"
	"github.com/xww2652008969/wbot/client"
	"github.com/xww2652008969/wbot/client/Message"
)

func main() {
	config := client.Clientconfig{
		Wsurl:      "ws://127.0.0.1",
		Wspost:     "3001",
		Wsheader:   nil,
		Clienthttp: "http://127.0.0.1:4000",
	}
	c, err := client.Create(config)
	if err != nil {
		panic(err)
	}
	var f client.Eventfunc
	f.Event = append(f.Event, client.Group)
	f.Func = append(f.Func, te1())
	f.Event = append(f.Event, client.Private)
	f.Func = append(f.Func, te2())
	c.EvebtFun = append(c.EvebtFun, f)
	c.Run()
	select {}

}
func te1() Message.MessageEvent {
	return func(message Message.Message) {
		fmt.Println("执行了 群聊")
		message.AddText("aa").AddText("aaa")
		message.SendGroupMsg(message.GroupId)
		message.AddText("第二次")
		message.SendGroupMsg(message.GroupId, false)
		message.AddText("清空了").SendGroupMsg(message.GroupId)
	}
}
func te2() Message.MessageEvent {
	return func(message Message.Message) {
		fmt.Print("执行了 私聊")
	}
}
