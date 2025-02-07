package main

import (
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
	c.RegisterPrivateHandle(Te())
	c.Run()
	select {}

}
func Te() client.Event {
	return func(c client.Client, message client.Message) {
		c.AddText("qqqqq")
		c.SendPrivateMsg(message.UserId)
	}
}
func Te2() client.Event {
	return func(c client.Client, message client.Message) {
		c.AddText("www")
		c.SendGroupMsg(message.GroupId)
	}
}
